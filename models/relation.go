package models

type Relation struct {
	UserId       string `bson:"userid" json:"userid,omitempty"`
	UserRelation string `bson:"userrelation" json:"userRelation,omitempty"`
}
