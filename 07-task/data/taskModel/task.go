package taskModel

import (
	"encoding/json"
	"fmt"
	"go.etcd.io/bbolt"
	"task/behelper"
	"task/bolt"
)

type Task struct {
	ID        int    `json:"id"`
	Details   string `json:"details"`
	Completed bool   `json:"completed"`
}

var tasksBucketName = []byte("tasks")

func CreateTask(taskDTO *Task) error {
	db, err := bolt.Connection()
	if err != nil {
		return err
	}
	err = db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(tasksBucketName)

		id, err := b.NextSequence()
		if err != nil {
			return err
		}

		taskDTO.ID = int(id)

		buf, err := json.Marshal(&taskDTO)
		if err != nil {
			return err
		}

		return b.Put(behelper.Itob(taskDTO.ID), buf)
	})
	if err != nil {
		return err
	}
	return nil
}

func ListTasks(bol bool) ([]*Task, error) {
	db, err := bolt.Connection()
	if err != nil {
		return nil, err
	}
	var tasks []*Task
	err = db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(tasksBucketName)
		return b.ForEach(func(_, v []byte) error {
			var task Task
			if err := json.Unmarshal(v, &task); err != nil {
				return err
			}
			tasks = append(tasks, &task)
			return nil
		})

	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func CompleteTask(taskDTO *Task) error {
	db, err := bolt.Connection()
	if err != nil {
		return err
	}
	return db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(tasksBucketName)

		task := b.Get(behelper.Itob(taskDTO.ID))
		if task == nil {
			return fmt.Errorf("taskModel not found with ID=%d", taskDTO.ID)
		}
		if err := json.Unmarshal(task, taskDTO); err != nil {
			return err
		}
		taskDTO.Completed = true
		task, err := json.Marshal(&taskDTO)
		if err != nil {
			return err
		}
		return b.Put(behelper.Itob(taskDTO.ID), task)
	})
}

func RemoveTask(taskDTO *Task) error {
	db, err := bolt.Connection()
	if err != nil {
		return err
	}
	return db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(tasksBucketName)

		task := b.Get(behelper.Itob(taskDTO.ID))
		if task == nil {
			return fmt.Errorf("taskModel not found with ID=%d", taskDTO.ID)
		}
		if err := json.Unmarshal(task, taskDTO); err != nil {
			return err
		}
		return b.DeleteBucket(behelper.Itob(taskDTO.ID))
	})
}
