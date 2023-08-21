package model

type ErrorResponse struct {
	Message string  `json:"message"`
	Title   *string `json:"title"`
}

type GetSkillsResponse struct {
	SkillsWithClasses []SkillWithClasses `json:"skillsWithClasses"`
	SkillsPerLvl      []SkillsPerLvlInfo `json:"skillsPerLvl"`
}
