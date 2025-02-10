package collections

import (
	"API_Config/internal/models\"
	resources "API_Config/internal/repositories/collections"
	"github.com/sirupsen/logrus"
)

func GetAllResources() ([]models.Resources, error) {
	var err error
	resources, err := resources.GetAllResources()
	if err != nil {
		logrus.Errorf("error retrieving resources : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return resources, nil
}
