package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"strings"
)

// map
var estateFeature_Map map[string]string = map[string]string{
	"最上階" : "saijoukai",
	"防犯カメラ" : "bouhan",
	"ウォークインクローゼット" : "walkin",
	"ワンルーム" : "oneroom",
	"ルーフバルコニー付" : "roofval",
	"エアコン付き" : "aircon",
	"駐輪場あり" : "chuurin",
	"プロパンガス" : "propan",
	"駐車場あり" : "chuushajou",
	"防音室" : "bouonn",
	"追い焚き風呂" : "oidaki",
	"オートロック" : "autolock",
	"即入居可" : "sokunyukyo",
	"IHコンロ" : "ihconro",
	"敷地内駐車場" : "sikichinai",
	"トランクルーム" : "tranckroom",
	"角部屋" : "kadobeya",
	"カスタマイズ可" : "customizable",
	"DIY可" : "diyok",
	"ロフト" : "roft",
	"シューズボックス" : "shoesbox",
	"インターネット無料" : "internetfree",
	"地下室" : "chikashitu",
	"敷地内ゴミ置場" : "shikichinaiGomiokiba",
	"管理人有り" : "kanrininari",
	"宅配ボックス" : "takuhaibokusu",
	"ルームシェア可" : "roomshare",
	"セキュリティ会社加入済" : "securitycomp",
	"メゾネット" : "mezonet",
	"女性限定" : "women",
	"バイク置場あり" : "baikuokiba",
	"エレベーター" : "eleveter",
	"ペット相談可" : "petokkamo",
	"洗面所独立" : "senmenjo",
	"都市ガス" : "toshigus",
	"浴室乾燥機" : "yokushitu",
	"インターネット接続可" : "internetokok",
	"テレビ・通信" : "televisionnetwork",
	"専用庭" : "niwa",
	"システムキッチン" : "systemkitchen",
	"高齢者歓迎" : "koureisha",
	"ケーブルテレビ" : "cabletv",
	"床下収納" : "yukashita",
	"バス・トイレ別" : "bustoiletbetsu",
	"駐車場2台以上" : "chuusha2ijou",
	"楽器相談可" : "gakisoudanok",
	"フローリング" : "flooring",
	"オール電化" : "alldenka",
	"TVモニタ付きインタホン" : "tvmonitorinter",
	"デザイナーズ物件" : "designers",
}

//Estate 物件
type Estate2 struct {
	ID          int64   `db:"id" json:"id"`

	Saijoukai bool  `db:"saijoukai"`
	Bouhan bool  `db:"bouhan"`
	Walkin bool  `db:"walkin"`
	Oneroom bool  `db:"oneroom"`
	Roofval bool  `db:"roofval"`
	Aircon bool  `db:"aircon"`
	Chuurin bool  `db:"chuurin"`
	Propan bool  `db:"propan"`
	Chuushajou bool  `db:"chuushajou"`
	Bouonn bool  `db:"bouonn"`
	Oidaki bool  `db:"oidaki"`
	Autolock bool  `db:"autolock"`
	Sokunyukyo bool  `db:"sokunyukyo"`
	Ihconro bool  `db:"ihconro"`
	Sikichinai bool  `db:"sikichinai"`
	Tranckroom bool  `db:"tranckroom"`
	Kadobeya bool  `db:"kadobeya"`
	Customizable bool  `db:"customizable"`
	Diyok bool  `db:"diyok"`
	Roft bool  `db:"roft"`
	Shoesbox bool  `db:"shoesbox"`
	Internetfree bool  `db:"internetfree"`
	Chikashitu bool  `db:"chikashitu"`
	Shikichinaigomiokiba bool  `db:"shikichinaiGomiokiba"`
	Kanrininari bool  `db:"kanrininari"`
	Takuhaibokusu bool  `db:"takuhaibokusu"`
	Roomshare bool  `db:"roomshare"`
	Securitycomp bool  `db:"securitycomp"`
	Mezonet bool  `db:"mezonet"`
	Women bool  `db:"women"`
	Baikuokiba bool  `db:"baikuokiba"`
	Eleveter bool  `db:"eleveter"`
	Petokkamo bool  `db:"petokkamo"`
	Senmenjo bool  `db:"senmenjo"`
	Toshigus bool  `db:"toshigus"`
	Yokushitu bool  `db:"yokushitu"`
	Internetokok bool  `db:"internetokok"`
	Televisionnetwork bool  `db:"televisionnetwork"`
	Niwa bool  `db:"niwa"`
	Systemkitchen bool  `db:"systemkitchen"`
	Koureisha bool  `db:"koureisha"`
	Cabletv bool  `db:"cabletv"`
	Yukashita bool  `db:"yukashita"`
	Bustoiletbetsu bool  `db:"bustoiletbetsu"`
	Chuusha2Ijou bool  `db:"chuusha2ijou"`
	Gakisoudanok bool  `db:"gakisoudanok"`
	Flooring bool  `db:"flooring"`
	Alldenka bool  `db:"alldenka"`
	Tvmonitorinter bool  `db:"tvmonitorinter"`
	Designers bool  `db:"designers"`

	Thumbnail   string  `db:"thumbnail" json:"thumbnail"`
	Name        string  `db:"name" json:"name"`
	Description string  `db:"description" json:"description"`
	Latitude    float64 `db:"latitude" json:"latitude"`
	Longitude   float64 `db:"longitude" json:"longitude"`
	Address     string  `db:"address" json:"address"`
	Rent        int64   `db:"rent" json:"rent"`
	DoorHeight  int64   `db:"door_height" json:"doorHeight"`
	DoorWidth   int64   `db:"door_width" json:"doorWidth"`
	Features    string  `db:"features" json:"features"`
	Popularity  int64   `db:"popularity" json:"-"`
}

func makeSetTextForFeature(feature string) string {

	var textSplitedList = strings.Split(feature, ",")
	var sql = ` SET `

	var setCondition = []string{}

	for _, feature_i := range textSplitedList {
		fieldName := estateFeature_Map[feature_i]
		if len(fieldName) == 0 {
			println("error make set text for feature! : ", feature_i)
			break
		}
		// sql +=  fmt.Sprintf("%s = 1,\n", fieldName)
		setCondition = append(setCondition, fmt.Sprintf("%s = 1", fieldName))
	}
	if len(setCondition) == 0 {
		return ""
	}

	sql += strings.Join(setCondition, ",")

	return sql
}

func makeInsertTextForEstate(t Estate) string {
	var sql = `INSERT INTO isuumo.estate `

	madeSqlSet := makeSetTextForFeature(t.Features)
	sql += madeSqlSet


	atoiId := strconv.FormatInt(t.ID, 10)
	sql += ` id = ` + atoiId + ", "
	sql += ` thumbnail =  ` + t.Thumbnail + ", "
	sql += ` name = ` + t.Name + ", "
	sql += ` description = ` + t.Description + ", "
	sql += ` latitude = ` + strconv.FormatFloat(t.Latitude, 'f', -1, 64) + ", "
	sql += ` longitude = ` + strconv.FormatFloat(t.Longitude, 'f', -1, 64) + ", "
	sql += ` address = ` + t.Address + ", "
	sql += ` rent = ` + strconv.FormatInt(t.Rent, 10) + ", "
	sql += ` door_height = ` + strconv.FormatInt(t.DoorHeight, 10) + ", "
	sql += ` door_width = ` + strconv.FormatInt(t.DoorWidth, 10) + ", "
	sql += ` features = ` + t.Features + ", "
	sql += ` popularity = ` + strconv.FormatInt(t.Popularity, 10)

	sql += ` ; `

	return sql
}

func convertEstate2ToEstate(t Estate2) Estate {
	return Estate{
		ID: t.ID,
		Name: t.Name,
		Description: t.Description,
		Thumbnail: t.Thumbnail,
		Address: t.Address,
		Latitude: t.Latitude,
		Longitude: t.Longitude,
		Rent: t.Rent,
		DoorHeight: t.DoorHeight,
		DoorWidth: t.DoorWidth,
		Features: t.Features,
		Popularity: t.Popularity,

	}
}

func migrationEstate() {
	tx, err := db.Beginx()
	if err != nil {
		println("db making failed transaction")
	}
	defer tx.Rollback()

	sql := `SELECT * FROM estate`
	estate2 := []Estate2{}
	err = tx.Select(&estate2, sql)
	println("ALL FETCH SELECT estate num: ", len(estate2))
	if err != nil {
		println ("migration select failed!")
	}

	sql =  `UPDATE estate `
	for _, v := range estate2 {
		madeStr := makeSetTextForFeature(v.Features)
		if len(madeStr) == 0 {
			continue
		}
		sql2 := sql + madeStr
		sql2 += ` WHERE id = ?;`
		
		tx.Exec(sql2, v.ID)
		println("      update sql in loop : ", sql2, " and v: ", fmt.Sprintf("%v{}", v))
	}
	println("update sql is for migration : ", sql)
	if err := tx.Commit(); err != nil {
		println("TX COMMIT FAILED!")
	}
	println("tx commit ended! in migration.. s")
}


func postEstate2(c echo.Context) error {
	header, err := c.FormFile("estates")
	if err != nil {
		c.Logger().Errorf("failed to get form file: %v", err)
		return c.NoContent(http.StatusBadRequest)
	}
	f, err := header.Open()
	if err != nil {
		c.Logger().Errorf("failed to open form file: %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}
	defer f.Close()
	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		c.Logger().Errorf("failed to read csv: %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	tx, err := db.Begin()
	if err != nil {
		c.Logger().Errorf("failed to begin tx: %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}
	defer tx.Rollback()
	for _, row := range records {
		rm := RecordMapper{Record: row}
		id := rm.NextInt()
		name := rm.NextString()
		description := rm.NextString()
		thumbnail := rm.NextString()
		address := rm.NextString()
		latitude := rm.NextFloat()
		longitude := rm.NextFloat()
		rent := rm.NextInt()
		doorHeight := rm.NextInt()
		doorWidth := rm.NextInt()
		features := rm.NextString()
		popularity := rm.NextInt()
		var estateVal = Estate{
			ID: int64(id),
			Name: name,
			Description: description,
			Thumbnail: thumbnail,
			Address: address,
			Latitude: latitude,
			Longitude: longitude,
			Rent: int64(rent),
			DoorHeight: int64(doorHeight),
			DoorWidth: int64(doorWidth),
			Features: features,
			Popularity: int64(popularity),
		}
		if err := rm.Err(); err != nil {
			c.Logger().Errorf("failed to read record: %v", err)
			return c.NoContent(http.StatusBadRequest)
		}

		createdSql := makeInsertTextForEstate(estateVal)

		//_, err := tx.Exec("INSERT INTO estate(id, name, description, thumbnail, address, latitude, longitude, rent, door_height, door_width, features, popularity) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)", id, name, description, thumbnail, address, latitude, longitude, rent, doorHeight, doorWidth, features, popularity)
		_, err := tx.Exec(createdSql)
		if err != nil {
			c.Logger().Errorf("failed to insert estate: %v", err)
			return c.NoContent(http.StatusInternalServerError)
		}
	}
	if err := tx.Commit(); err != nil {
		c.Logger().Errorf("failed to commit tx: %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusCreated)
}


func dealWithSearchFeatureEstate2(features string, conditions *[]string) {
	if features != "" {
		for _, fea := range strings.Split(features, ",") {
			fieldInMysql := estateFeature_Map[fea]
			cond := fieldInMysql + `  =  true`

			//*conditions = append(*conditions, "features like concat('%', ?, '%')")
			*conditions = append(*conditions, cond)

		}
	}
}

func searchEstates2(c echo.Context) error {
	conditions := make([]string, 0)
	params := make([]interface{}, 0)

	if c.QueryParam("doorHeightRangeId") != "" {
		doorHeight, err := getRange(estateSearchCondition.DoorHeight, c.QueryParam("doorHeightRangeId"))
		if err != nil {
			c.Echo().Logger.Infof("doorHeightRangeID invalid, %v : %v", c.QueryParam("doorHeightRangeId"), err)
			return c.NoContent(http.StatusBadRequest)
		}

		if doorHeight.Min != -1 {
			conditions = append(conditions, "door_height >= ?")
			params = append(params, doorHeight.Min)
		}
		if doorHeight.Max != -1 {
			conditions = append(conditions, "door_height < ?")
			params = append(params, doorHeight.Max)
		}
	}

	if c.QueryParam("doorWidthRangeId") != "" {
		doorWidth, err := getRange(estateSearchCondition.DoorWidth, c.QueryParam("doorWidthRangeId"))
		if err != nil {
			c.Echo().Logger.Infof("doorWidthRangeID invalid, %v : %v", c.QueryParam("doorWidthRangeId"), err)
			return c.NoContent(http.StatusBadRequest)
		}

		if doorWidth.Min != -1 {
			conditions = append(conditions, "door_width >= ?")
			params = append(params, doorWidth.Min)
		}
		if doorWidth.Max != -1 {
			conditions = append(conditions, "door_width < ?")
			params = append(params, doorWidth.Max)
		}
	}

	if c.QueryParam("rentRangeId") != "" {
		estateRent, err := getRange(estateSearchCondition.Rent, c.QueryParam("rentRangeId"))
		if err != nil {
			c.Echo().Logger.Infof("rentRangeID invalid, %v : %v", c.QueryParam("rentRangeId"), err)
			return c.NoContent(http.StatusBadRequest)
		}

		if estateRent.Min != -1 {
			conditions = append(conditions, "rent >= ?")
			params = append(params, estateRent.Min)
		}
		if estateRent.Max != -1 {
			conditions = append(conditions, "rent < ?")
			params = append(params, estateRent.Max)
		}
	}

	dealWithSearchFeatureEstate2(c.QueryParam("features"), &conditions)
	//if c.QueryParam("features") != "" {
	//	for _, f := range strings.Split(c.QueryParam("features"), ",") {
	//		conditions = append(conditions, "features like concat('%', ?, '%')")
	//		params = append(params, f)
	//	}
	//}

	if len(conditions) == 0 {
		c.Echo().Logger.Infof("searchEstates search condition not found")
		return c.NoContent(http.StatusBadRequest)
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		c.Logger().Infof("Invalid format page parameter : %v", err)
		return c.NoContent(http.StatusBadRequest)
	}

	perPage, err := strconv.Atoi(c.QueryParam("perPage"))
	if err != nil {
		c.Logger().Infof("Invalid format perPage parameter : %v", err)
		return c.NoContent(http.StatusBadRequest)
	}

	searchQuery := "SELECT * FROM estate WHERE "
	countQuery := "SELECT COUNT(*) FROM estate WHERE "
	searchCondition := strings.Join(conditions, " AND ")
	limitOffset := " ORDER BY popularity DESC, id ASC LIMIT ? OFFSET ?"

	var res EstateSearchResponse
	err = db.Get(&res.Count, countQuery+searchCondition, params...)
	if err != nil {
		c.Logger().Errorf("searchEstates DB execution error : %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	estates := []Estate{}
	estates2 := []Estate2{}
	params = append(params, perPage, page*perPage)
	err = db.Select(&estates2, searchQuery+searchCondition+limitOffset, params...)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusOK, EstateSearchResponse{Count: 0, Estates: []Estate{}})
		}
		c.Logger().Errorf("searchEstates DB execution error : %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	for _, est := range estates2 {
		estates = append(estates, convertEstate2ToEstate(est))
	}

	res.Estates = estates

	return c.JSON(http.StatusOK, res)
}
