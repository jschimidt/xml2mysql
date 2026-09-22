package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Job struct {
		Name string `yaml:"name"`
	} `yaml:"job"`
	Input struct {
		Directory string `yaml:"directory"`
		Pattern   string `yaml:"pattern"`
	} `yaml:"input"`
	Output struct {
		Directory string `yaml:"directory"`
	} `yaml:"output"`
	Analysis struct {
		Workers       int `yaml:"workers"`
		DistinctLimit int `yaml:"distinct_limit"`
		SampleLimit   int `yaml:"sample_limit"`
	} `yaml:"analysis"`
	Target struct {
		Database     string `yaml:"database"`
		Version      string `yaml:"version"`
		DatabaseName string `yaml:"database_name"`
		Engine       string `yaml:"engine"`
		Charset      string `yaml:"charset"`
	} `yaml:"target"`
}

func Load(path string) (Config, error) {
	var c Config
	b, e := os.ReadFile(path)
	if e != nil {
		return c, e
	}
	if e = yaml.Unmarshal(b, &c); e != nil {
		return c, e
	}
	if c.Job.Name == "" {
		c.Job.Name = "default"
	}
	if c.Input.Directory == "" {
		return c, fmt.Errorf("input.directory is required")
	}
	if c.Input.Pattern == "" {
		c.Input.Pattern = "*.xml"
	}
	if c.Output.Directory == "" {
		return c, fmt.Errorf("output.directory is required")
	}
	if c.Analysis.Workers < 1 {
		c.Analysis.Workers = 1
	}
	if c.Analysis.DistinctLimit < 1 {
		c.Analysis.DistinctLimit = 100
	}
	if c.Analysis.SampleLimit < 1 {
		c.Analysis.SampleLimit = 10
	}
	if c.Target.Database == "" {
		c.Target.Database = "mysql"
	}
	if c.Target.Version == "" {
		c.Target.Version = "8.4"
	}
	if c.Target.DatabaseName == "" {
		c.Target.DatabaseName = c.Job.Name + "_sistema_ids_bkp"
	}
	if c.Target.Engine == "" {
		c.Target.Engine = "InnoDB"
	}
	if c.Target.Charset == "" {
		c.Target.Charset = "utf8mb4"
	}
	c.Output.Directory = filepath.Join(c.Output.Directory, c.Job.Name)
	return c, nil
}
