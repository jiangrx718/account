package tags

import (
	"account/internal/common"
	"account/internal/service"

	//"account/internal/dao"
	"context"
)

type TagRep struct {
	TagId    string `json:"tag_id"`
	TeamId   string `json:"team_id"`
	TagName  string `json:"tag_name"`
	TagColor string `json:"tag_color"`
	TagType  int32  `json:"tag_type"`
}

func (s *TagService) TagCreate(ctx context.Context, teamId, tagName, tagColor string, tagType int32) (common.ServiceResult, error) {
	var (
		//logObj = log.SugarContext(ctx)
		result = service.NewServiceResult()
	)
	//
	//tagData := model.Tags{
	//	TagId:    utils.GenUUID(),
	//	TeamId:   teamId,
	//	TagName:  tagName,
	//	TagColor: tagColor,
	//	TagType:  &tagType,
	//}
	//
	//// 验证标签是否已经存在
	//tagDao := dao.Tags
	//count, err := tagDao.Where(
	//	tagDao.TeamId.Eq(teamId),
	//	tagDao.TagName.Eq(tagName),
	//	tagDao.TagType.Eq(tagType),
	//).Count()
	//if err != nil {
	//	logObj.Errorw("TagCreate count error", "teamId", teamId, "tagName", tagName, "error", err)
	//	return result, err
	//}
	//
	//if count > 0 {
	//	result.SetError(&common.ServiceError{
	//		Code:    400,
	//		Message: "标签已存在",
	//	})
	//	return result, nil
	//}
	//
	//if createErr := tagDao.Create(&tagData); createErr != nil {
	//	logObj.Errorw("TagCreate error", "tag", tagData, "error", createErr)
	//	return result, createErr
	//}
	//
	//result.Data = TagRep{
	//	TagId:    tagData.TagId,
	//	TeamId:   tagData.TeamId,
	//	TagName:  tagData.TagName,
	//	TagColor: tagData.TagColor,
	//	TagType:  *tagData.TagType,
	//}
	result.SetMessage("操作成功")
	return result, nil
}
