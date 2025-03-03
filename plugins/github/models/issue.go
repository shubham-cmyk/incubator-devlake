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
	"time"
)

type GithubIssue struct {
	ConnectionId    uint64 `gorm:"primaryKey"`
	GithubId        int    `gorm:"primaryKey"`
	RepoId          int    `gorm:"index"`
	Number          int    `gorm:"index;comment:Used in API requests ex. api/repo/1/issue/<THIS_NUMBER>"`
	State           string `gorm:"type:varchar(255)"`
	Title           string
	Body            string
	Priority        string `gorm:"type:varchar(255)"`
	Type            string `gorm:"type:varchar(100)"`
	StdType         string `gorm:"type:varchar(100)"`
	AuthorId        int
	AuthorName      string `gorm:"type:varchar(255)"`
	AssigneeId      int
	AssigneeName    string `gorm:"type:varchar(255)"`
	MilestoneId     int    `gorm:"index"`
	LeadTimeMinutes uint
	Url             string `gorm:"type:varchar(255)"`
	ClosedAt        *time.Time
	GithubCreatedAt time.Time
	GithubUpdatedAt time.Time `gorm:"index"`
	Severity        string    `gorm:"type:varchar(255)"`
	Component       string    `gorm:"type:varchar(255)"`
	common.NoPKModel
}

func (GithubIssue) TableName() string {
	return "_tool_github_issues"
}
