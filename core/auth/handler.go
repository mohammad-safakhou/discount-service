package auth

import (
	"discount-service/utils"
	"encoding/json"
	"errors"
	"github.com/spf13/viper"
	"log"
	"strings"

	"github.com/labstack/echo"
)


// GetAuthentication check and get auth from given token
func getAuthentication(config *viper.Viper, token string, accountID string) (*Identity, error) {
	if token == "" || len(token) < 7 {
		return nil, utils.StandardHttpErrorResponse{Code: 401}
	}
	if token[:7] == "bearer " {
		token = token[7:]
	}
	// get jwt data
	jwtDataStr := strings.Split(token, ".")[1]
	bytes, err := base64Decode(jwtDataStr)
	if err != nil {
		log.Println(err)
		return nil, utils.StandardHttpErrorResponse{Code: 401}
	}
	properties := make(map[string]interface{})
	err = json.Unmarshal(bytes, &properties)
	if err != nil {
		log.Println(err)
		return nil, utils.StandardHttpErrorResponse{Code: 401}
	}
	// find secret key of application
	secretKey := ""
	// clientID := ""
	if app, ok := properties["app"]; ok {
		applicationID := app.(string)
		_, secretKey, err = GetApplicationAuthKeys(config, applicationID)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("Invalid JWT")
	}
	claims, err := ValidateJWT(token, secretKey)
	if err != nil {
		log.Println(err)
		return nil, utils.StandardHttpErrorResponse{Code: 401}
	}
	serviceAvailable := false
	var service IdentityService
	if svc, ok := claims["svc"]; ok {
		if services, ok := svc.(map[string]interface{}); ok {
			if _service, ok := services[config.GetString("env.service_name")]; ok {
				if serviceMap, ok := _service.(map[string]interface{}); ok {
					serviceAvailable = true
					perm, ok := serviceMap["perm"]
					var permissions int64 = 0
					if ok {
						permissions = int64(perm.(float64))
					}
					service = IdentityService{
						Permissions: permissions,
					}
				}
			}
		}
	}

	// nameID := claims["nameid"].(string)
	userID := claims["id"].(string)

	var rolesMap map[string]bool
	if _roles, ok := claims["role"]; ok {
		rolesMap = make(map[string]bool)
		if roles, ok := _roles.([]interface{}); ok {
			for _, role := range roles {
				rolesMap[role.(string)] = true
			}
		}
	}
	if !serviceAvailable {
		// check if is rostam
		if _, ok := rolesMap["rostam"]; !ok {
			return nil, utils.StandardHttpErrorResponse{Code: 401}
		}
	}
	merchantRoles := map[string][]string{}
	if _merchantRoles, ok := claims["merchant_roles"]; ok {
		if merchantRolesMapFace, ok := _merchantRoles.(map[string]interface{}); ok {
			for merchantRolesKey, _mRoles := range merchantRolesMapFace {
				if mRoles, ok := _mRoles.([]interface{}); ok {
					merchantRoles[merchantRolesKey] = []string{}
					for _, _role := range mRoles {
						merchantRoles[merchantRolesKey] = append(merchantRoles[merchantRolesKey], _role.(string))
					}
				}
			}
		}
	}

	var AppID string
	AppID = claims["app"].(string)
	if accountID != "" && userID != accountID {
		accountAccessGranted := false
		exists := false
		if rolesMap != nil {
			if _, ok := rolesMap["zeus"]; ok {
				exists = true
			} else if _, ok := rolesMap["rostam"]; ok {
				exists = true
			}
		}
		if exists {
			accountAccessGranted = true
		}
		if !accountAccessGranted {
			if _, ok := merchantRoles[accountID]; ok {
				accountAccessGranted = true
			} else {
				return nil, utils.StandardHttpErrorResponse{Code: 403}
			}
		} else {
			return nil, utils.StandardHttpErrorResponse{Code: 403}
		}
	}
	identity := Identity{Token: token, ID: userID, Roles: rolesMap, MerchantRoles: merchantRoles, Service: service, AppId: AppID}
	return &identity, nil
}

// Authenticate current request if is valid, otherwise returns nil
func (ac *AuthenticateContext) Authenticate(c *echo.Context, accountID string, roles ...string) (*Identity, error) {
	token := (*c).Request().Header.Get("Authorization")
	identity, err := getAuthentication(ac.VConfig, token, accountID)
	if err == nil && len(roles) > 0 {
		exists := false
		for _, role := range roles {
			if _, ok := identity.Roles[role]; ok {
				exists = true
				break
			}
		}
		if !exists {
			return nil, utils.StandardHttpErrorResponse{Code: 401}
		}
	}
	return identity, err
}
