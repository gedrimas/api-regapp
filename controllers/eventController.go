package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"api-regapp/database"
	"api-regapp/helpers"
	"api-regapp/models"
	
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gedrimas/response"

)

type createEventReq struct {
	Event       string             `json:"event" validate:"required"`
	Start       time.Time          `json:"start" validate:"required"`
	End         time.Time          `json:"end" validate:"required"`
	Description string             `json:"description"`
	Trainer     string             `json:"trainer"`
	Places      string             `json:"places" validate:"required"`
	Visitors    int                `json:"visitors"`
	Day         string             `json:"day"`
}

var validate = validator.New()
var eventCollection *mongo.Collection = database.OpenCollection(database.Client, "event")

func CreateEvent() gin.HandlerFunc {
	return func(c *gin.Context) {

		sendor := response.NewSendor(c)

		if err := helpers.VerifyUserType(c, "ADMIN"); err != nil {
			badRequest := response.GetResponse("badRequest")
			badRequest.SetData(err.Error())
			sendor.Send(badRequest)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var req createEventReq
		
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status":  http.StatusBadRequest,
				"Message": "error",
				"Data":    map[string]interface{}{"data": err.Error()}})
			return
		}

		eventId := primitive.NewObjectID()
		companyId := c.GetString("company_id") 
		adminId := c.GetString("user_id")

		newEvent := models.Event{
			Id:          eventId,
			Event_id:    eventId.Hex(),
			Event:       req.Event,
			Start:       req.Start,
			End:         req.End,
			Description: req.Description,
			Trainer:     req.Trainer,
			Places:      req.Places,
			Visitors:    req.Visitors,
			Day:         req.Day,
			Company_id:  companyId,
			Admin_id:    adminId,
		}

		if validRequestError := validate.Struct(&newEvent); validRequestError != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Status":  http.StatusBadRequest,
				"Message": "error",
				"Data":    map[string]interface{}{"data": validRequestError.Error()}})
			return
		}

		result, err := eventCollection.InsertOne(ctx, newEvent)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Status":  http.StatusInternalServerError,
				"Message": "error",
				"Data":    map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"Status":  http.StatusCreated,
			"Message": "error",
			"Data":    map[string]interface{}{"data": result},
		})
	}
}

func GetEvent() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		//var event models.Event
		defer cancel()
		companyId := c.GetString("company_id")
		fmt.Println("COMPANY", companyId)


		events, err := eventCollection.Find(ctx, bson.D{{"company_id", companyId}})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Status":  http.StatusInternalServerError,
				"Message": "error",
				"Data":    map[string]interface{}{"data": err.Error()}})
			return
		}

		var results []bson.M
		if err = events.All(ctx, &results); err != nil {
			log.Fatal("EVENTS", err)
		}


		c.JSON(http.StatusOK, gin.H{
			"Status":  http.StatusOK,
			"Message": "success",
			"Data":    map[string]interface{}{"data": results}})
	}
}
