package parser

import "encoding/xml"

type KaddressImageId struct {
	XMLName xml.Name `xml:"addressImageId,omitempty" json:"addressImageId,omitempty"`
	Text    string
}

type KaddressOutside struct {
	XMLName xml.Name `xml:"addressOutside,omitempty" json:"addressOutside,omitempty"`
	Text    string
}

type KaddressOutsideImageId struct {
	XMLName xml.Name `xml:"addressOutsideImageId,omitempty" json:"addressOutsideImageId,omitempty"`
	Text    string
}

type KassignmentDate struct {
	XMLName xml.Name `xml:"assignmentDate,omitempty" json:"assignmentDate,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type KchangeCause struct {
	XMLName xml.Name `xml:"changeCause,omitempty" json:"changeCause,omitempty"`
	Text    string
}

type KchangeDate struct {
	XMLName xml.Name `xml:"changeDate,omitempty" json:"changeDate,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type KcityCode struct {
	XMLName xml.Name `xml:"cityCode,omitempty" json:"cityCode,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type KcityName struct {
	XMLName xml.Name `xml:"cityName,omitempty" json:"cityName,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type KcloseCause struct {
	XMLName xml.Name `xml:"closeCause,omitempty" json:"closeCause,omitempty"`
	Text    string
}

type KcloseDate struct {
	XMLName xml.Name `xml:"closeDate,omitempty" json:"closeDate,omitempty"`
	Text    string
}

type KcorporateNumber struct {
	XMLName xml.Name `xml:"corporateNumber,omitempty" json:"corporateNumber,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type Kcorporation struct {
	XMLName                   xml.Name                   `xml:"corporation,omitempty" json:"corporation,omitempty"`
	KaddressImageId           *KaddressImageId           `xml:"addressImageId,omitempty" json:"addressImageId,omitempty"`
	KaddressOutside           *KaddressOutside           `xml:"addressOutside,omitempty" json:"addressOutside,omitempty"`
	KaddressOutsideImageId    *KaddressOutsideImageId    `xml:"addressOutsideImageId,omitempty" json:"addressOutsideImageId,omitempty"`
	KassignmentDate           *KassignmentDate           `xml:"assignmentDate,omitempty" json:"assignmentDate,omitempty"`
	KchangeCause              *KchangeCause              `xml:"changeCause,omitempty" json:"changeCause,omitempty"`
	KchangeDate               *KchangeDate               `xml:"changeDate,omitempty" json:"changeDate,omitempty"`
	KcityCode                 *KcityCode                 `xml:"cityCode,omitempty" json:"cityCode,omitempty"`
	KcityName                 *KcityName                 `xml:"cityName,omitempty" json:"cityName,omitempty"`
	KcloseCause               *KcloseCause               `xml:"closeCause,omitempty" json:"closeCause,omitempty"`
	KcloseDate                *KcloseDate                `xml:"closeDate,omitempty" json:"closeDate,omitempty"`
	KcorporateNumber          *KcorporateNumber          `xml:"corporateNumber,omitempty" json:"corporateNumber,omitempty"`
	Kcorrect                  *Kcorrect                  `xml:"correct,omitempty" json:"correct,omitempty"`
	KenAddressOutside         *KenAddressOutside         `xml:"enAddressOutside,omitempty" json:"enAddressOutside,omitempty"`
	KenCityName               *KenCityName               `xml:"enCityName,omitempty" json:"enCityName,omitempty"`
	KenName                   *KenName                   `xml:"enName,omitempty" json:"enName,omitempty"`
	KenPrefectureName         *KenPrefectureName         `xml:"enPrefectureName,omitempty" json:"enPrefectureName,omitempty"`
	Kfurigana                 *Kfurigana                 `xml:"furigana,omitempty" json:"furigana,omitempty"`
	Khihyoji                  *Khihyoji                  `xml:"hihyoji,omitempty" json:"hihyoji,omitempty"`
	Kkind                     *Kkind                     `xml:"kind,omitempty" json:"kind,omitempty"`
	Klatest                   *Klatest                   `xml:"latest,omitempty" json:"latest,omitempty"`
	Kname                     *Kname                     `xml:"name,omitempty" json:"name,omitempty"`
	KnameImageId              *KnameImageId              `xml:"nameImageId,omitempty" json:"nameImageId,omitempty"`
	KpostCode                 *KpostCode                 `xml:"postCode,omitempty" json:"postCode,omitempty"`
	KprefectureCode           *KprefectureCode           `xml:"prefectureCode,omitempty" json:"prefectureCode,omitempty"`
	KprefectureName           *KprefectureName           `xml:"prefectureName,omitempty" json:"prefectureName,omitempty"`
	Kprocess                  *Kprocess                  `xml:"process,omitempty" json:"process,omitempty"`
	KsequenceNumber           *KsequenceNumber           `xml:"sequenceNumber,omitempty" json:"sequenceNumber,omitempty"`
	KstreetNumber             *KstreetNumber             `xml:"streetNumber,omitempty" json:"streetNumber,omitempty"`
	KsuccessorCorporateNumber *KsuccessorCorporateNumber `xml:"successorCorporateNumber,omitempty" json:"successorCorporateNumber,omitempty"`
	KupdateDate               *KupdateDate               `xml:"updateDate,omitempty" json:"updateDate,omitempty"`
}

type Kcorporations struct {
	XMLName         xml.Name         `xml:"corporations,omitempty" json:"corporations,omitempty"`
	Kcorporation    []*Kcorporation  `xml:"corporation,omitempty" json:"corporation,omitempty"`
	Kcount          *Kcount          `xml:"count,omitempty" json:"count,omitempty"`
	KdivideNumber   *KdivideNumber   `xml:"divideNumber,omitempty" json:"divideNumber,omitempty"`
	KdivideSize     *KdivideSize     `xml:"divideSize,omitempty" json:"divideSize,omitempty"`
	KlastUpdateDate *KlastUpdateDate `xml:"lastUpdateDate,omitempty" json:"lastUpdateDate,omitempty"`
}

type Kcorrect struct {
	XMLName xml.Name `xml:"correct,omitempty" json:"correct,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type Kcount struct {
	XMLName xml.Name `xml:"count,omitempty" json:"count,omitempty"`
	Text    int      `xml:",chardata" json:",omitempty"`
}

type KdivideNumber struct {
	XMLName xml.Name `xml:"divideNumber,omitempty" json:"divideNumber,omitempty"`
	Text    int      `xml:",chardata" json:",omitempty"`
}

type KdivideSize struct {
	XMLName xml.Name `xml:"divideSize,omitempty" json:"divideSize,omitempty"`
	Flag    bool     `xml:",chardata" json:",omitempty"`
}

type KenAddressOutside struct {
	XMLName xml.Name `xml:"enAddressOutside,omitempty" json:"enAddressOutside,omitempty"`
	Text    string
}

type KenCityName struct {
	XMLName xml.Name `xml:"enCityName,omitempty" json:"enCityName,omitempty"`
	Text    string
}

type KenName struct {
	XMLName xml.Name `xml:"enName,omitempty" json:"enName,omitempty"`
	Text    string
}

type KenPrefectureName struct {
	XMLName xml.Name `xml:"enPrefectureName,omitempty" json:"enPrefectureName,omitempty"`
	Text    string
}

type Kfurigana struct {
	XMLName xml.Name `xml:"furigana,omitempty" json:"furigana,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type Khihyoji struct {
	XMLName xml.Name `xml:"hihyoji,omitempty" json:"hihyoji,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type Kkind struct {
	XMLName xml.Name `xml:"kind,omitempty" json:"kind,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type KlastUpdateDate struct {
	XMLName xml.Name `xml:"lastUpdateDate,omitempty" json:"lastUpdateDate,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type Klatest struct {
	XMLName xml.Name `xml:"latest,omitempty" json:"latest,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type Kname struct {
	XMLName xml.Name `xml:"name,omitempty" json:"name,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type KnameImageId struct {
	XMLName xml.Name `xml:"nameImageId,omitempty" json:"nameImageId,omitempty"`
	Text    string
}

type KpostCode struct {
	XMLName xml.Name `xml:"postCode,omitempty" json:"postCode,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type KprefectureCode struct {
	XMLName xml.Name `xml:"prefectureCode,omitempty" json:"prefectureCode,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type KprefectureName struct {
	XMLName xml.Name `xml:"prefectureName,omitempty" json:"prefectureName,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type Kprocess struct {
	XMLName xml.Name `xml:"process,omitempty" json:"process,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type KsequenceNumber struct {
	XMLName xml.Name `xml:"sequenceNumber,omitempty" json:"sequenceNumber,omitempty"`
	Number  int      `xml:",chardata" json:",omitempty"`
}

type KstreetNumber struct {
	XMLName xml.Name `xml:"streetNumber,omitempty" json:"streetNumber,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}

type KsuccessorCorporateNumber struct {
	XMLName xml.Name `xml:"successorCorporateNumber,omitempty" json:"successorCorporateNumber,omitempty"`
	Text    string
}

type KupdateDate struct {
	XMLName xml.Name `xml:"updateDate,omitempty" json:"updateDate,omitempty"`
	Text    string   `xml:",chardata" json:",omitempty"`
}
