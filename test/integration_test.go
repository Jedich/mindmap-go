package test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mindmap-go/app/models"
	"mindmap-go/app/repository"
	"mindmap-go/app/services"
	"mindmap-go/internal/database"
	"testing"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite")
}

var (
	db   *database.Database
	repo repository.CardRepository
	card *models.Card
)

var _ = BeforeSuite(func() {
	conn, err := gorm.Open(mysql.Open("root:11@tcp(localhost:3306)/mindmap_test?parseTime=true"), &gorm.Config{})
	Expect(err).ShouldNot(HaveOccurred())

	db = &database.Database{Connection: conn}
	err = db.Migrate()
	Expect(err).ShouldNot(HaveOccurred())

	repo = repository.NewCardRepository(db)
})

var _ = AfterSuite(func() {
	err := db.Connection.Exec("DROP TABLE accounts, cards, files, maps, users;").Error
	Expect(err).ShouldNot(HaveOccurred())
	db.CloseConnection()
})

var _ = Describe("Integration Tests", func() {
	Context("Card Repository", func() {
		It("should be empty", func() {
			l, err := repo.GetCardsByMapID(0)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(l).Should(BeEmpty())
		})

		It("should create", func() {
			userRepo := repository.NewUserRepository(db)
			accRepo := repository.NewAccountRepository(db)
			hashedPwd, err := services.NewUserService(userRepo, accRepo).Hash("pwd")
			Expect(err).ShouldNot(HaveOccurred())

			user := &models.User{
				Account: models.Account{
					Username:     "username",
					Email:        "email@a.com",
					PasswordHash: hashedPwd,
				},
			}
			err = userRepo.CreateUser(user)
			Expect(err).ShouldNot(HaveOccurred())

			mapRepo := repository.NewMapRepository(db)
			mapp := &models.Map{
				CreatorID: user.ID,
				Cards:     nil,
			}
			err = mapRepo.CreateMap(mapp)
			Expect(err).ShouldNot(HaveOccurred())
			cardToCreate := &models.Card{
				ParentID:  nil,
				CreatorID: user.ID,
				MapID:     mapp.ID,
			}
			err = repo.CreateCard(cardToCreate)
			Expect(err).ShouldNot(HaveOccurred())
			card = cardToCreate
		})

		It("should not be empty", func() {
			l, err := repo.GetCardsByMapID(card.MapID)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(l).ShouldNot(BeEmpty())
		})

		It("should update", func() {
			err := repo.UpdateCard(card)
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})
