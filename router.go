package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

//routes yönlendirmerler
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Controller.go dosyası
type UserController struct {
	Session *mgo.Session
	Uye     Uye
}
type DilController struct {
	Session *mgo.Session
	D       Dil
	Di      DilInfo
}
type MesajController struct {
	Session *mgo.Session
	M       Mesaj
	MI      MesajInfo
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	return router
}

var routes = Routes{
	Route{"UserGet", "POST", "/user/get", uc.Get},
	Route{"UserCreate", "POST", "/user/create", uc.Create},
	Route{"UserUpdate", "POST", "/user/update", uc.Update},
	Route{"UserDelete", "DELETE", "/user/delete", uc.Remove},
	Route{"UserGetAll", "GET", "/user/all", uc.GetAll},
	Route{"UserLogin", "POST", "/user/login", uc.Login},
	//Route{"UserDeneme", "POST", "/user/deneme", uc.Deneme},
	Route{"UserGetAllUyeler", "POST", "/user/uyeler", uc.GetAllUyeler},
	Route{"UserEmailKontrol", "POST", "/user/email", uc.EmailKontrol},
	Route{"UserArkadasEkle", "POST", "/user/arkadas/ekle", uc.ArkadasEkle},
	Route{"UserArkadasOnayla", "POST", "/user/arkadas/onayla", uc.ArkadaslikOnayla},
	Route{"UserArkadasCikar", "POST", "/user/arkadas/delete", uc.ArkadasCikar},
	Route{"UserArkadasAll", "POST", "/user/arkadas/all", uc.ArkadasAll},
	Route{"UserArkadasIstekler", "POST", "/user/arkadas/istekler", uc.ArkadasIstekler},
	Route{"UserArkadasArama", "POST", "/user/arkadas/arama", uc.ArkadasArama},
	Route{"UserAyarlarGet", "POST", "/user/ayar", uc.AyarlarGet},
	Route{"UserAyarlarUpdate", "POST", "/user/ayar/updaye", uc.AyarlarUpdate},

	Route{"Mesajlar", "POST", "/mesaj/all", mc.MesajlarimiGetir},
	Route{"MesajlarSon", "POST", "/mesaj", mc.SonMesajiGetir},
	Route{"MesajlarGonder", "POST", "/mesaj/gonder", mc.Gonder},

	Route{"DilCreate", "POST", "/dil/create", dc.Create},
	Route{"DilGetAll", "GET", "/dil/all", dc.GetAll},
	Route{"DilUpdate", "PUT", "/dil/update", dc.Update},
	Route{"DilDelete", "DELETE", "/dil/delete", dc.Delete},
}
