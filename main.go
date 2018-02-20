package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	//"strconv"

	//	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var uc = NewUserController(getSession())
var dc = NewDilController(getSession())
var mc = NewMesajController(getSession())

func main() {
	r := mux.NewRouter().StrictSlash(true)
	for _, a := range routes {
		r.HandleFunc(a.Pattern, a.HandlerFunc).
			Methods(a.Method)
	}

	http.ListenAndServe("localhost:8080", handlers.CORS()(r))
}

//mongodb session acama
func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}

/*func (uc UserController) Deneme(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)

	u := User{}
	json.NewDecoder(r.Body).Decode(&u)
	// Add an Id
	fmt.Println("istek:", u)

	session := getSession()
	/*c := session.DB("test").C("uyes")
	var results []Uye
	errr := c.Find(nil).All(&results)
	if errr != nil {
		panic(errr)
	}

	d := Deneme{}
	/*for i := range results {
		d.Arkadaslar[results[i].Id.String()] = results[i]
	}
	c1 := session.DB("test").C("deneme")
	err := c1.Find(bson.M{"_id": u.Id}).One(&d)
	if _, ok := d.Arkadaslar[bson.ObjectIdHex(u.Email).String()]; ok {
		delete(d.Arkadaslar, bson.ObjectIdHex(u.Email).String())
	}
	err = c1.Update(bson.M{"_id": d.Id}, d)
	if err != nil {
		log.Fatal(err)
	}
	uj, err := json.Marshal(d)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("%s\n", uj)
	fmt.Fprintf(w, "%s", uj)
}
*/

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{Session: s, Uye: Uye{}}
}
func NewDilController(s *mgo.Session) *DilController {
	return &DilController{Session: s, D: Dil{}, Di: DilInfo{}}
}
func NewMesajController(s *mgo.Session) *MesajController {
	return &MesajController{Session: s, M: Mesaj{}, MI: MesajInfo{}}
}

//Mesajlar ile alakalı
func (mc MesajController) Insert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)

	m := Mesaj{}
	json.NewDecoder(r.Body).Decode(&m)
	//m.Insert(m)
}

func (mc MesajController) MesajlarimiGetir(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)

	m := Mesaj{}
	json.NewDecoder(r.Body).Decode(&m)
	fmt.Printf("%s \n", m.Id)

	a := mc.M.GetByUyeId(m.Id)
	uj, _ := json.Marshal(a)

	fmt.Printf("%s \n", uj)
	fmt.Fprintf(w, "%s", uj)
}
func (mc MesajController) SonMesajiGetir(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)

	m := Mesaj{}
	json.NewDecoder(r.Body).Decode(&m)
	m = mc.M.Get(m.Id)
	uj, _ := json.Marshal(m.Mesajlar)

	fmt.Printf("%s \n", uj)
	fmt.Fprintf(w, "%s", uj)
}
func (mc MesajController) Gonder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)

	m := Mesaj{}
	json.NewDecoder(r.Body).Decode(&m)

	if m.Id.String() == "" {
		mc.M.Create(m)
	} else {
		for _, value := range m.Mesajlar {
			mc.M.Insert(value)
		}
		//mc.M.Insert()
	}

	uj, _ := json.Marshal(m)

	fmt.Printf("%s \n", uj)
	fmt.Fprintf(w, "%s", uj)
}

//Mesajlar bitiş
///Dilller alakalı controller
func (dc DilController) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
	d := Dil{}
	json.NewDecoder(r.Body).Decode(&d)
	// Add an Id
	dc.D.Create(d)

}

func (dc DilController) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(200)
	d := Dil{}
	json.NewDecoder(r.Body).Decode(&d)
	// Add an Id
	dc.D.Update(d)

}
func (dc DilController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Access-Control-Allow-Credentials", "true")
	//Access-Control-Allow-Origin: http://127.0.0.1:8080
	//w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	if origin := r.Header.Get("Origin"); origin == "http://localhost:8080" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	w.WriteHeader(200)
	d := Dil{}
	json.NewDecoder(r.Body).Decode(&d)
	// Add an Id
	//dc.D.Delete(d)
	uj, _ := json.Marshal(d)

	fmt.Printf("%s \n", uj)
	fmt.Fprintf(w, "%s", uj)
}

func (dc DilController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)

	session := getSession()
	c := session.DB("test").C("diller")
	result := []Dil{}
	c.Find(nil).All(&result)
	uj, err := json.Marshal(result)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Fprintf(w, "%s", uj)
	fmt.Printf("%s\n", uj)
	// Add an Id
}

///dilller alakalı bitiş
//uye getir
func (uc UserController) Get(w http.ResponseWriter, r *http.Request) {
	u := User{}
	json.NewDecoder(r.Body).Decode(&u)
	//session := getSession()
	//c := session.DB("test").C("uyes")
	uye := uc.Uye.Get(u.Id)
	//err := c.Find(bson.M{"Id": bson.ObjectIdHex(id)}).One(&u)
	///if err != nil {
	//	log.Fatal(err)
	//}
	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(uye)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(uj)
	fmt.Printf("%S", uj)
}

//uye oluştur
func (uc UserController) Create(w http.ResponseWriter, r *http.Request) {
	u := User{}
	json.NewDecoder(r.Body).Decode(&u)
	//u.Id = bson.NewObjectId()
	uj, _ := json.Marshal(u)
	/*
		Uye struct {
			Id         bson.ObjectId `json:"id" bson:"_id"`
			Isım       string        `json:"adsoyad" bson:"adsoyad"`
			Email      string        `json:"email" bson:"email"`
			Sifre      string        `json:"adsoyad" bson:"adsoyad"`
			Profil     Profil        `json:"profil" bson:"profil"`
			Arkadaslar []Profil      `json:"arkadaslar" bson:"arkadaslar"`
			Mesajlar   []Mesaj       `json:"mesajlar" bson:"mesajlar"`
			Dil        Dil           `json:"dil" bson:"dil"`
			Durum      bool          `json:"durum" bson:"durum"`
			Ayarlar    Ayar          `json:"ayar" bson:"ayar"`
			Aktif      bool          `json:"aktif" bson:"aktif"`
		}
		Profil struct {
			Id       bson.ObjectId `json:"id" bson:"_id"`
			AdSoyad  string        `json:"adsoyad" bson:"adsoyad"`
			Yas      int           `json:"yas" bson:"yas"`
			Meslek   string        `json:"meslek" bson:"meslek"`
			Aciklama string        `json:"aciklama" bson:"aciklama"`
			Resimurl string        `json:"resimurl" bson:"resimurl"`
			//diller   []DilInfo			`json:"diller" bson:"diller"`
			Anadil      Dil      `json:"anadil" bson:"anadil"`
			Ogr_diller  []Dil    `json:"odiller" bson:"odiller"`
			Biln_diller []Dil    `json:"bdiller" bson:"bdiller"`
			Ulke        string   `json:"ulke" bson:"ulke"`
			Sehir       string   `json:"sehir" bson:"sehir"`
			Hobiler     string   `json:"hobiler" bson:"hobiler"`
			Arkadaslar  []Profil `json:"arkadaslar" bson:"arkadaslar"`
			Hakkinda    string   `json:"hakkinda" bson:"hakkinda"`
			Amac        string   `json:"amac" bson:"amac"`
			Cinsiyet    string   `json:"cinsiyet" bson:"cinsiyet"`
			Dogumtarihi string   `json:"dtarihi" bson:"dtarihi"`
		}*/
	dil := Dil{Ad: u.Dil, Resimurl: ""}
	p := Profil{AdSoyad: u.Isım, Ulke: u.Ulke, Sehir: u.Il, Dogumtarihi: u.Dtarihi, Cinsiyet: u.Cinsiyet, Resimurl: u.Presim, Anadil: dil, Ogr_diller: make(map[string]Dil), Biln_diller: make(map[string]Dil), Arkadaslar: make(map[string]Profil)}
	uy := Uye{Email: u.Email, Sifre: u.Sifre, Aktif: false, Arkadaslar: make(map[string]Uye), Ayarlar: Ayar{}, Isım: u.Isım, Profil: p, Mesajlar: make(map[string]Mesaj), Dil: dil, Durum: false, Istekler: make(map[string]Arkadaslik)}
	var id = bson.NewObjectId()
	fmt.Printf("%s\n", id)
	uy.Id = id
	p.Id = id
	uc.Uye.CreateUye(uy)
	uc.Uye.CreateProfil(p)
	uju, _ := json.Marshal(uy)
	ujp, _ := json.Marshal(p)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
	fmt.Printf("\n%s", uju)
	fmt.Printf("\n%s", ujp)
}
func (uc UserController) Update(w http.ResponseWriter, r *http.Request) {
	u := Profil{}
	// Populate the user data
	json.NewDecoder(r.Body).Decode(&u)
	uye := uc.Uye.Get(u.Id)
	uye.Profil = u
	uye.Isım = u.AdSoyad
	uye.Dil.Ad = u.Anadil.Ad
	uc.Uye.Update(uye)
	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)
	fmt.Printf("\n%s", uj)
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

//uye kaldır
func (uc UserController) Remove(w http.ResponseWriter, r *http.Request) {
	// Grab id
	vars := mux.Vars(r)
	id := vars["id"]
	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}
	// Grab id
	oid := bson.ObjectIdHex(id)
	// Remove user
	session := getSession()
	c := session.DB("test").C("peoples")
	if err := c.RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}
	// Write status
	w.WriteHeader(200)
}

//deneme amaclı
func (uc UserController) GetAll(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
	c := getSession().DB("test").C("uyes")
	var results []Uye
	errr := c.Find(nil).All(&results)
	if errr != nil {
		panic(errr)
	}

	ujj, err := json.Marshal(results)
	if err != nil {
		fmt.Print("err")
	}
	fmt.Fprintf(w, "%s", ujj)

	fmt.Printf("%s\n", ujj)
}
func (uc UserController) GetAllUyeler(w http.ResponseWriter, r *http.Request) {
	u := User{}
	json.NewDecoder(r.Body).Decode(&u)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
	result := uc.Uye.GetAll()
	user := uc.Uye.Get(u.Id)
	leng := len(result)
	tmp := make([]Uye, leng)
	a := 0
	for i := 0; i < len(result); i++ {
		for _, value := range user.Arkadaslar {
			if value.Id != result[i].Id {
				tmp[a] = value
				a++
			}
		}
	}

	ujj, err := json.Marshal(tmp)
	if err != nil {
		fmt.Print("err")
	}
	fmt.Fprintf(w, "%s", ujj)

	fmt.Printf("%s\n", ujj)
}
func (uc UserController) Login(w http.ResponseWriter, r *http.Request) {
	u := Uye{}
	json.NewDecoder(r.Body).Decode(&u)
	u = uc.Uye.Login(u.Email, u.Sifre)
	uj, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(uj)
	fmt.Printf("%s", uj)
}
func (uc UserController) EmailKontrol(w http.ResponseWriter, r *http.Request) {
	u := User{}
	json.NewDecoder(r.Body).Decode(&u)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
	c := getSession().DB("test").C("uyes")
	//	var results []Uye
	//SearchReading(bson.M{"k": key, "t": { $gte: start, $lte: end } }, limit)
	count, errr := c.Find(bson.M{"email": u.Email}).Count()
	if errr != nil {
		log.Fatal(errr)
	}

	fmt.Fprintf(w, "%d", count)

	fmt.Printf("%d\n", count)
}

func (uc UserController) ArkadasEkle(w http.ResponseWriter, r *http.Request) {
	u := User{}

	json.NewDecoder(r.Body).Decode(&u)
	fmt.Printf("istek gelen:%s email:%s\n", u.Id.String(), u.Email)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)

	uc.Uye.ArkadaslikEkle(bson.ObjectIdHex(u.Email), u.Id)

}
func (uc UserController) ArkadasCikar(w http.ResponseWriter, r *http.Request) {
	u := User{}
	json.NewDecoder(r.Body).Decode(&u)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)

	//	uye := uc.Uye.Get(bson.ObjectIdHex(u.Email))
	//	arkadas := uc.Uye.GetProfil(u.Id)

	uc.Uye.ArkadaslikCikar(bson.ObjectIdHex(u.Email), u.Id)

	/*	i := len(uye.Arkadaslar)

			//	id := arkadas.Id
			fmt.Println("uzunluk dizi:", i)
			if _, ok := uye.Arkadaslar[u.Id.String()]; ok {
				delete(uye.Arkadaslar, u.Id.String())

				fmt.Println("kaldırıldı. Durum:", ok)
			} else {
				fmt.Println("Zaten yokmuş")
			}

			uc.Uye.Update(uye)

			uj, err := json.Marshal(arkadas)
			if err != nil {
				log.Fatal(err)
			}

		w.Write(uj)
		fmt.Printf("\n%s\n", uj)*/

}
func (uc UserController) ArkadaslikGetir(w http.ResponseWriter, r *http.Request) {
	u := User{}
	json.NewDecoder(r.Body).Decode(&u)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)

	tmp := uc.Uye.ArkadaslikIstekleri(u.Id)

	uj, err := json.Marshal(tmp)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(uj)
	fmt.Printf("\n%s\n", uj)

}
func (uc UserController) ArkadaslikOnayla(w http.ResponseWriter, r *http.Request) {
	u := User{}
	json.NewDecoder(r.Body).Decode(&u)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
	fmt.Println(u)
	uc.Uye.ArkadaslikOnayla(u.Id, bson.ObjectIdHex(u.Email))

	/*uj, err := json.Marshal(tmp)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(uj)
	fmt.Printf("\n%s\n", uj)*/

}

func (uc UserController) ArkadasIstekler(w http.ResponseWriter, r *http.Request) {
	u := User{}
	json.NewDecoder(r.Body).Decode(&u)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)

	tmp := uc.Uye.Get(u.Id)

	uj, err := json.Marshal(tmp.Istekler)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(uj)
	fmt.Printf("\n%s\n", uj)

}
func (uc UserController) ArkadasAll(w http.ResponseWriter, r *http.Request) {
	u := User{}
	json.NewDecoder(r.Body).Decode(&u)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)

	uye := uc.Uye.Get(u.Id)
	//	id := arkadas.Id
	uj, err := json.Marshal(uye.Arkadaslar)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(uj)
	fmt.Printf("\n%s\n", uj)

}
func (uc UserController) ArkadasArama(w http.ResponseWriter, r *http.Request) {
	u := Sorgu{}
	json.NewDecoder(r.Body).Decode(&u)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)

	//uye := uc.uye.Get(u.Id)
	//	id := arkadas.Id
	result := uc.Uye.GetAllbyQuery(u)
	uj, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(uj)
	fmt.Printf("\n%s\n", uj)

}
func (uc UserController) AyarlarUpdate(w http.ResponseWriter, r *http.Request) {
	u := Ayar{}
	json.NewDecoder(r.Body).Decode(&u)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)

	//uye := uc.uye.Get(u.Id)
	//	id := arkadas.Id
	result := uc.Uye.AyarlarUpdate(u.Id, u)
	uj, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(uj)
	fmt.Printf("\n%s\n", uj)

}
func (uc UserController) AyarlarGet(w http.ResponseWriter, r *http.Request) {
	u := Ayar{}
	json.NewDecoder(r.Body).Decode(&u)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)

	result := uc.Uye.AyarlarGet(u.Id)
	uj, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(uj)
	fmt.Printf("\n%s\n", uj)

}
