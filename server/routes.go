package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"log"
	"strconv"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"server": "1",
	})
}

func CreateSegment(c *gin.Context) {
	var segment Segmentlist
	c.BindJSON(&segment)
	query := fmt.Sprintf("INSERT INTO segmentlist (slug) VALUES ('%s')", segment.Slug)
	_, err := DB.Exec(query)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"server": -1,
		})
	} else {
		c.JSON(200, gin.H{
			"server": 1,
		})
	}
}

func DeleteSegment(c *gin.Context) {
	var segment Segmentlist
	c.BindJSON(&segment)
	query := fmt.Sprintf("DELETE FROM segmentlist WHERE slug='%s'", segment.Slug)
	_, err := DB.Exec(query)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"server": -1,
		})
	} else {
		c.JSON(200, gin.H{
			"server": 1,
		})
	}
}

func AddUserToSegment(c *gin.Context) {
	// check if user exists
	var upd Updater
	c.BindJSON(&upd)
	query := fmt.Sprintf("SELECT segments FROM userss WHERE user_id=%s", strconv.Itoa(upd.UserID))
	rows, err := DB.Query(query)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"server": -1,
		})
	} else {
		counter := 0
		for rows.Next() {
			counter += 1
			break
		}
		if counter == 0 {
			// user does not exist
			// creating user, adding to segments
			// check if there are no elements from SlugsToDelete in SlugsToAdd:
			var toDel []int32
			for _, el := range upd.SlugsToDelete {
				for _, sl := range upd.SlugsToAdd {
					if el == sl {
						toDel = append(toDel, el)
					}
				}
			}
			if len(toDel) == len(upd.SlugsToAdd) {
				upd.SlugsToAdd = nil
			} else {
				for _, el := range toDel {
					upd.SlugsToAdd = append(upd.SlugsToAdd[:el], upd.SlugsToAdd[el+1:]...)
				}
			}

			// no adding to DB

			if len(upd.SlugsToAdd) == 0 {
				query = fmt.Sprintf("INSERT INTO userss (segments) values ('{}')")
				_, err = DB.Exec(query)
				if err != nil {
					log.Println(err)
					c.JSON(500, gin.H{
						"server": -1,
					})
				} else {
					c.JSON(200, gin.H{
						"server": 1,
						"msg":    "User was created (no slugs).",
					})
				}
			} else {
				_, err = DB.Exec("INSERT INTO userss (segments) values ($1)", pq.Array(upd.SlugsToAdd))
				if err != nil {
					log.Println(err)
					c.JSON(500, gin.H{
						"server": -1,
					})
				} else {
					c.JSON(200, gin.H{
						"server": 1,
						"msg":    "User was created and updated.",
					})
				}
			}

		} else {
			// user exists
			var existingSlugs []int32
			var user User
			query = fmt.Sprintf("SELECT segments FROM userss WHERE user_id=%s", strconv.Itoa(upd.UserID))
			row := DB.QueryRow(query)
			row.Scan(&user.UserID, &user.Segments)
			//existingSlugs = PQtoArray(string(user.Segments))
			existingSlugs = user.Segments
			// perform similar check from above
			// check if there are no elements from SlugsToDelete in SlugsToAdd:
			var toDel []int32
			for _, el := range upd.SlugsToDelete {
				for _, sl := range upd.SlugsToAdd {
					if el == sl {
						toDel = append(toDel, el)
					}
				}
			}
			if len(toDel) == len(upd.SlugsToAdd) {
				upd.SlugsToAdd = nil
			} else {
				for _, el := range toDel {
					upd.SlugsToAdd = append(upd.SlugsToAdd[:el], upd.SlugsToAdd[el+1:]...)
				}
			}
			// deleting from user's slugs
			var toDel2 []int32
			for _, el := range existingSlugs {
				for _, sl := range upd.SlugsToDelete {
					if el == sl {
						toDel2 = append(toDel2, sl)
					}
				}
			}
			// deleting
			if len(toDel2) == len(existingSlugs) {
				existingSlugs = nil
			} else {
				for _, el := range toDel2 {
					existingSlugs = append(existingSlugs[:el], existingSlugs[el+1:]...)
				}
			}

			// now we have clean existingSlugs and SlugsToAdd

			for _, el := range upd.SlugsToAdd {
				existingSlugs = append(existingSlugs, el)
			}

			if len(existingSlugs) == 0 {
				query = fmt.Sprintf("UPDATE userss SET segments='{}' WHERE user_id=%s", strconv.Itoa(upd.UserID))
				_, err = DB.Exec(query)
				if err != nil {
					log.Println(err)
					c.JSON(500, gin.H{
						"server": -1,
					})
				} else {
					c.JSON(200, gin.H{
						"server": 1,
						"msg":    "Existing user was updated. No slugs now.",
					})
				}
			} else {
				//query = fmt.Sprintf("UPDATE userss SET segments='{%s}' WHERE user_id=%s", ArrayToPQ(existingSlugs), strconv.Itoa(upd.UserID))
				_, err = DB.Exec("UPDATE userss SET segments=$1 WHERE user_id=$2", pq.Array(existingSlugs), strconv.Itoa(upd.UserID))
				if err != nil {
					log.Println(err)
					c.JSON(500, gin.H{
						"server": -1,
					})
				} else {
					c.JSON(200, gin.H{
						"server": 1,
						"msg":    "Existing user was updated.",
					})
				}
			}

		}
	}
}

func GetUserSegments(c *gin.Context) {
	id := c.Query("user_id")
	query := fmt.Sprintf("SELECT segments FROM userss WHERE user_id=%d", First(strconv.Atoi(id)))
	row := DB.QueryRow(query)
	var seg string
	row.Scan(&seg)
	if seg == "" {
		c.JSON(200, gin.H{
			"server": 0,
			"msg":    "No such user.",
		})
	} else {
		c.JSON(200, gin.H{
			"server": 1,
			"msg":    seg,
		})
	}
}

func GetAllUsers(c *gin.Context) {
	var userlist []User
	query := fmt.Sprintf("SELECT * FROM userss")
	rows, err := DB.Query(query)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"server": -1,
		})
	} else {
		for rows.Next() {
			var usr User
			rows.Scan(&usr.UserID, &usr.Segments)
			userlist = append(userlist, usr)
		}
		c.JSON(200, gin.H{
			"users": userlist,
		})
	}
}
