package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"net/http"
	"os"
)

// HelmValues represents the structure of the Helm values file
type HelmValues struct {
	Image struct {
		Repository string `yaml:"repository"`
		Tag        string `yaml:"tag"`
		PullPolicy string `yaml:"pullPolicy"`
	} `yaml:"image"`
	Service struct {
		Type string `yaml:"type"`
		Port int    `yaml:"port"`
	} `yaml:"service"`
}

func main() {
	router := gin.Default()

	router.POST("/generate-helm-values", func(c *gin.Context) {
		var values HelmValues
		if err := c.ShouldBindJSON(&values); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		yamlData, err := yaml.Marshal(&values)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate YAML"})
			return
		}

		filePath := "values.yaml"
		if err := os.WriteFile(filePath, yamlData, 0644); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write file"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Helm values file generated", "filePath": filePath})
	})

	router.Run(":8080")
}
