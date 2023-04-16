package api

import (
	"encoding/json"
	"fmt"
	"time"
)

func (a *HttpApi) FetchBookMetaInfo(workId int, userToken string) (*BookMetaInfo, error) {
	path := fmt.Sprintf("v1/work/%d/details", workId)
	body, err := makeRequest(path, userToken)
	if err != nil {
		return nil, err
	}

	bmi := &BookMetaInfo{}
	json.Unmarshal(body, bmi)

	return bmi, nil
}

type ChapterMeta struct {
	ID                   int       `json:"id"`
	WorkID               int       `json:"workId"`
	Title                string    `json:"title"`
	IsDraft              bool      `json:"isDraft"`
	SortOrder            int       `json:"sortOrder"`
	PublishTime          time.Time `json:"publishTime"`
	LastModificationTime time.Time `json:"lastModificationTime"`
	TextLength           int       `json:"textLength"`
	IsAvailable          bool      `json:"isAvailable"`
}

type BookMetaInfo struct {
	Chapters                []ChapterMeta `json:"chapters"`
	AllowDownloads          bool          `json:"allowDownloads"`
	DownloadErrorCode       string        `json:"downloadErrorCode"`
	DownloadErrorMessage    string        `json:"downloadErrorMessage"`
	PrivacyDownloads        string        `json:"privacyDownloads"`
	Annotation              string        `json:"annotation"`
	AuthorNotes             string        `json:"authorNotes"`
	AtRecommendation        bool          `json:"atRecommendation"`
	SeriesWorkIds           []int         `json:"seriesWorkIds"`
	SeriesWorkNumber        int           `json:"seriesWorkNumber"`
	ReviewCount             int           `json:"reviewCount"`
	Tags                    []string      `json:"tags"`
	OrderID                 any           `json:"orderId"`
	OrderStatus             string        `json:"orderStatus"`
	OrderStatusMessage      any           `json:"orderStatusMessage"`
	Contests                []any         `json:"contests"`
	GalleryImages           []any         `json:"galleryImages"`
	BooktrailerVideoURL     any           `json:"booktrailerVideoUrl"`
	IsExclusive             bool          `json:"isExclusive"`
	FreeChapterCount        int           `json:"freeChapterCount"`
	PromoFragment           bool          `json:"promoFragment"`
	LinkedWork              any           `json:"linkedWork"`
	ID                      int           `json:"id"`
	Title                   string        `json:"title"`
	CoverURL                string        `json:"coverUrl"`
	LastModificationTime    time.Time     `json:"lastModificationTime"`
	LastUpdateTime          time.Time     `json:"lastUpdateTime"`
	FinishTime              time.Time     `json:"finishTime"`
	IsFinished              bool          `json:"isFinished"`
	TextLength              int           `json:"textLength"`
	TextLengthLastRead      int           `json:"textLengthLastRead"`
	Price                   float64       `json:"price"`
	Discount                int           `json:"discount"`
	WorkForm                int           `json:"workForm"`
	Status                  string        `json:"status"`
	AuthorID                int           `json:"authorId"`
	AuthorFIO               string        `json:"authorFIO"`
	AuthorUserName          string        `json:"authorUserName"`
	OriginalAuthor          string        `json:"originalAuthor"`
	Translator              string        `json:"translator"`
	Reciter                 string        `json:"reciter"`
	CoAuthorID              int           `json:"coAuthorId"`
	CoAuthorFIO             string        `json:"coAuthorFIO"`
	CoAuthorUserName        string        `json:"coAuthorUserName"`
	CoAuthorConfirmed       bool          `json:"coAuthorConfirmed"`
	SecondCoAuthorID        int           `json:"secondCoAuthorId"`
	SecondCoAuthorFIO       string        `json:"secondCoAuthorFIO"`
	SecondCoAuthorUserName  string        `json:"secondCoAuthorUserName"`
	SecondCoAuthorConfirmed bool          `json:"secondCoAuthorConfirmed"`
	IsPurchased             bool          `json:"isPurchased"`
	UserLikeID              any           `json:"userLikeId"`
	LastReadTime            any           `json:"lastReadTime"`
	LastChapterID           any           `json:"lastChapterId"`
	LastChapterProgress     float64       `json:"lastChapterProgress"`
	LikeCount               int           `json:"likeCount"`
	CommentCount            int           `json:"commentCount"`
	RewardCount             int           `json:"rewardCount"`
	RewardsEnabled          bool          `json:"rewardsEnabled"`
	InLibraryState          string        `json:"inLibraryState"`
	AddedToLibraryTime      time.Time     `json:"addedToLibraryTime"`
	PrivacyDisplay          string        `json:"privacyDisplay"`
	State                   string        `json:"state"`
	IsDraft                 bool          `json:"isDraft"`
	EnableRedLine           bool          `json:"enableRedLine"`
	EnableTTS               bool          `json:"enableTTS"`
	AdultOnly               bool          `json:"adultOnly"`
	SeriesID                int           `json:"seriesId"`
	SeriesOrder             int           `json:"seriesOrder"`
	SeriesTitle             string        `json:"seriesTitle"`
	Afterword               any           `json:"afterword"`
	SeriesNextWorkID        int           `json:"seriesNextWorkId"`
	GenreID                 int           `json:"genreId"`
	FirstSubGenreID         int           `json:"firstSubGenreId"`
	SecondSubGenreID        int           `json:"secondSubGenreId"`
	Format                  string        `json:"format"`
}
