package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UnderController struct {
	Parking     ParkingRepository
	ParkingGate ParkingGateRepository
	Device      DeviceRepository
	Inside      InsideRepository
	ParamRepo   ParamRepository
	Event       NewEvent
}

func NewUnderController(parking ParkingRepository, parkingGate ParkingGateRepository, device DeviceRepository, inside InsideRepository, paramRepo ParamRepository) *UnderController {
	return &UnderController{
		Parking:     parking,
		ParkingGate: parkingGate,
		Device:      device,
		Inside:      inside,
		ParamRepo:   paramRepo,
		Event:       NewEvent{},
	}
}

func (ctrl *UnderController) MachineSeats(c *gin.Context) {
	log.Println("Получен GET запрос")
	c.HTML(http.StatusOK, "index.html", gin.H{"event": ctrl.Event})
}

func (ctrl *UnderController) Layout(c *gin.Context) {
	c.HTML(http.StatusOK, "layout.html", gin.H{"event": ctrl.Event})
}

func (ctrl *UnderController) Seats(c *gin.Context) {
	parkinData := ctrl.Parking.FindAllBy()
	c.HTML(http.StatusOK, "machine-seats.html", gin.H{
		"event":         ctrl.Event,
		"parking":       parkinData,
		"insertParking": InsertParking{},
	})
}

func (ctrl *UnderController) EditParking(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	parkingData := ctrl.Parking.FindAllById(id)
	c.HTML(http.StatusOK, "edit-parking.html", gin.H{
		"event":         ctrl.Event,
		"parking":       parkingData,
		"parkingGate":   ctrl.ParkingGate.GetAllGate(id),
		"parkingInside": ctrl.Inside.SelectAllInside(id),
	})
}

func (ctrl *UnderController) EditGate(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	idDev, _ := strconv.ParseInt(c.Param("idDev"), 10, 64)
	var hlParam HlParam
	var gate HlParkingGate
	c.HTML(http.StatusOK, "edit-gate.html", gin.H{
		"event":      ctrl.Event,
		"gate":       gate,
		"device":     ctrl.Device.GetAllBy(),
		"deviceById": ctrl.ParkingGate.GetDeviceById(idDev),
		"param":      hlParam,
	})
}

func (ctrl *UnderController) AddGate(c *gin.Context) {
	var gate HlParkingGate
	id, _ := strconv.Atoi(c.Param("id"))
	c.HTML(http.StatusOK, "add-gate.html", gin.H{
		"event":  ctrl.Event,
		"device": ctrl.Device.GetAllBy(),
	})
}

func (ctrl *UnderController) EditParkingImpl(c *gin.Context) {
	var hlParking HlParking
	if err := c.ShouldBind(&hlParking); err == nil {
		ctrl.Parking.UpdateParking(hlParking)
		ctrl.Event.SetEvent(true)
		c.Redirect(http.StatusFound, "/machine-seats")
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (ctrl *UnderController) DeleteParking(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	ctrl.Parking.DeleteHlParkingById(id)
	ctrl.Event.SetEvent(true)
	c.Redirect(http.StatusFound, "/machine-seats")
}

func (ctrl *UnderController) InsertParking(c *gin.Context) {
	var insertParking InsertParking
	if err := c.BindJSON(&insertParking); err == nil {
		ctrl.Parking.InsertParkingByName(insertParking.Name)
		ctrl.Event.SetEvent(true)
		c.Redirect(http.StatusFound, "/machine-seats")
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func (ctrl *UnderController) DeleteGate(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	idDev, _ := strconv.ParseInt(c.Param("idDev"), 10, 64)
	ctrl.ParkingGate.DeleteById(idDev)
	ctrl.Event.SetEvent(true)
	c.Redirect(http.StatusFound, "/edit-parking/"+strconv.FormatInt(id, 10))
}

func (ctrl *UnderController) SaveGateByID(c *gin.Context) {
	var gate HlParkingGate
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	idDev, _ := strconv.ParseInt(c.Param("idDev"), 10, 64)
	if err := c.ShouldBind(&gate); err == nil {
		ctrl.ParkingGate.UpdateGate(id, idDev, gate.IsEnter)
		ctrl.Event.SetEvent(true)
		c.Redirect(http.StatusFound, "/edit-parking/"+strconv.FormatInt(id, 10))
	}
}

func (ctrl *UnderController) SaveGate(c *gin.Context) {
	var gate HlParkingGate
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	name := c.PostForm("name")
	isEnter, _ := strconv.Atoi(c.PostForm("is_enter"))
	ctrl.ParkingGate.InsertGate(id, name, isEnter)
	ctrl.Event.SetEvent(true)
	c.Redirect(http.StatusFound, "/edit-parking/"+strconv.FormatInt(id, 10))
}

func (ctrl *UnderController) SaveGateDev(c *gin.Context) {
	id := c.Param("id")
	idDev := c.Param("idDev")
	// Uncomment for additional logic
	// ctrl.ParamRepo.AddParamDevice()
	c.Redirect(http.StatusFound, "/edit-parking/"+id)
}
