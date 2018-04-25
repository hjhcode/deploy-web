package managers

import (
	"strconv"

	"github.com/hjhcode/deploy-web/models"
)

func getMirrorName(mirrorId string) string {
	id, _ := strconv.ParseInt(mirrorId, 10, 64)
	mirror, err := models.Mirror{}.GetById(id)
	if err != nil {
		panic(err.Error())
	}
	if mirror == nil {
		return ""
	}

	return mirror.MirrorName
}

func getMirrorNameById(mirrorId int64) string {
	if mirrorId == 0 {
		return ""
	}
	mirror, err := models.Mirror{}.GetById(mirrorId)
	if err != nil {
		panic(err.Error())
	}
	if mirror == nil {
		return ""
	}

	return mirror.MirrorName
}
