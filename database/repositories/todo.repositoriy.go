package repositories

import (
	"fmt"

	"github.com/Marrquito/GoToDoAPP/database"
	"github.com/Marrquito/GoToDoAPP/database/models"
	"gorm.io/gorm"
)

type toDoRepository struct {
	database *gorm.DB
}

type ToDoRepository interface {
	Create(toDo models.ToDo) (models.ToDo, error)
	Read(id uint64) (*models.ToDo, error)
	List(options *ListToDoOptions) ([]models.ToDo, int64, error)
	Update(toDo models.ToDo) (models.ToDo, error)
	Delete(id uint64) error
}

type ListToDoOptions struct {
	Page   int
	Limit  int
	Order  string
	Offset int
	Load   *[]string
	Where  *models.ToDo
}

func NewToDoRepository() ToDoRepository {
	return &toDoRepository{
		database: database.Database,
	}
}

func (r *toDoRepository) Create(toDo models.ToDo) (models.ToDo, error) {
	var createdToDo models.ToDo

	db := r.database.Create(&toDo).Scan(&createdToDo)

	return createdToDo, db.Error
}

func (r *toDoRepository) Read(id uint64) (*models.ToDo, error) {
	var toDo models.ToDo

	db := r.database.Model(&models.ToDo{})
	db.First(&toDo, id)

	return &toDo, db.Error
}

func (r *toDoRepository) List(options *ListToDoOptions) ([]models.ToDo, int64, error) {
	var toDos []models.ToDo
	var total int64

	db := r.database.Model(&models.ToDo{})
	database.LoadRelations(db, options.Load)

	if options.Where != nil {
		db = db.Where(options.Where)
	}

	if options.Limit > 0 && options.Page > 0 {
		options.Offset = (options.Page - 1) * options.Limit
		db.Count(&total)
	}

	db.Limit(options.Limit).
		Offset(options.Offset).
		Order(fmt.Sprintf("id %s", options.Order)).
		Find(&toDos)

	return toDos, total, db.Error
}

func (r *toDoRepository) Update(toDo models.ToDo) (models.ToDo, error) {
	db := r.database.Save(&toDo)

	return toDo, db.Error
}

func (r *toDoRepository) Delete(id uint64) error {
	db := r.database.Delete(&models.ToDo{}, id)

	return db.Error
}