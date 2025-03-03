/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package models

import (
	"github.com/apache/incubator-devlake/models/common"
	"gorm.io/datatypes"
)

type GitlabTransformationRule struct {
	common.Model
	Name                 string            `gorm:"type:varchar(255);index:idx_name_gitlab,unique" validate:"required" mapstructure:"name" json:"name"`
	PrType               string            `mapstructure:"prType" json:"prType"`
	PrComponent          string            `mapstructure:"prComponent" json:"prComponent"`
	PrBodyClosePattern   string            `mapstructure:"prBodyClosePattern" json:"prBodyClosePattern"`
	IssueSeverity        string            `mapstructure:"issueSeverity" json:"issueSeverity"`
	IssuePriority        string            `mapstructure:"issuePriority" json:"issuePriority"`
	IssueComponent       string            `mapstructure:"issueComponent" json:"issueComponent"`
	IssueTypeBug         string            `mapstructure:"issueTypeBug" json:"issueTypeBug"`
	IssueTypeIncident    string            `mapstructure:"issueTypeIncident" json:"issueTypeIncident"`
	IssueTypeRequirement string            `mapstructure:"issueTypeRequirement" json:"issueTypeRequirement"`
	DeploymentPattern    string            `mapstructure:"deploymentPattern" json:"deploymentPattern"`
	ProductionPattern    string            `mapstructure:"productionPattern,omitempty" json:"productionPattern" gorm:"type:varchar(255)"`
	Refdiff              datatypes.JSONMap `mapstructure:"refdiff,omitempty" json:"refdiff" swaggertype:"object" format:"json"`
}

func (t GitlabTransformationRule) TableName() string {
	return "_tool_gitlab_transformation_rules"
}
