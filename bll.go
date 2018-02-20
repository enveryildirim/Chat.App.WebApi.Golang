package main

import (
	//	"fmt"
	//	"fmt"
	"log"
	"reflect"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Uye ile alakalı temel fonk lar
func (us Uye) CreateUye(u Uye) {
	session := getSession()
	c := session.DB("test").C("uyes")
	err := c.Insert(&u)
	if err != nil {
		log.Fatal(err)
	}
}
func (us Uye) CreateProfil(p Profil) {

	session := getSession()
	c := session.DB("test").C("profils")
	err := c.Insert(&p)
	if err != nil {
		log.Fatal(err)
	}
}

func (us Uye) Get(id bson.ObjectId) Uye {
	session := getSession()
	c := session.DB("test").C("uyes")
	u := Uye{}
	err := c.Find(bson.M{"_id": id}).One(&u)
	if err != nil {
		log.Fatal(err)
	}
	return u
}
func (us Uye) GetAllbyQuery(s Sorgu) []Uye {
	session := getSession()
	c := session.DB("test").C("uyes")
	var result []Uye
	/*db.uyes.find({
	    "isim": "enver",
	    "cinsiyet": "erkek",
	    "azyas": {
	        "$lt": 10
	    },
	    "cokyas": {
	        "$gt": 20
	    },
	    "online": 1,
	    "ulke": "turkiye",
	    "dil": "t\u00fcrk\u00e7e"
	});*/
	//bson.M{"profil.cinsiyet": s.Cinsiyet,"profil.durum": s.Online,"profil.ulke": s.Ulke,"profil.dil": s.Dil}
	err := c.Find(bson.M{"profil.cinsiyet": s.Cinsiyet, "durum": s.Online}).All(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
func (us Uye) GetAll() []Uye {
	session := getSession()
	c := session.DB("test").C("uyes")
	var result []Uye

	err := c.Find(nil).All(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
func (us Uye) Update(u Uye) {
	session := getSession()
	c := session.DB("test").C("uyes")
	err := c.Update(bson.M{"_id": u.Id}, u)
	if err != nil {
		log.Fatal(err)
	}
}
func (us Uye) Delete(id string) {

	oid := bson.ObjectIdHex(id)
	session := getSession()
	c := session.DB("test").C("uyes")
	if err := c.RemoveId(oid); err != nil {
		return
	}
}

//eksik düzeltilecek
func (us Uye) Login(email string, sifre string) Uye {
	session := getSession()
	c := session.DB("test").C("uyes")
	u := Uye{}
	err := c.Find(bson.M{"email": email, "sifre": sifre}).One(&u)
	if err != nil {
		log.Fatal(err)
	}
	return u

}

func (us Uye) Activate(id bson.ObjectId) {
	u := us.Get(id)
	u.Durum = true
	us.Update(u)
}

func EmailKontrol(email string) bool {
	session := getSession()
	c := session.DB("test").C("uyes")
	i, err := c.Find(bson.M{"email": email}).Count()
	if err != nil {
		log.Fatal(err)
	}
	return i == 0
}
func ArkadasContains(s []Profil, e Profil) bool {
	for _, a := range s {
		if reflect.DeepEqual(a, e) {
			return true
		}
	}
	return false
}
func ArkadasRemove(plist []Profil, p Profil) bool {
	if len(plist) < 1 {
		return false
	}
	for i := range plist {
		if reflect.DeepEqual(plist[i], p) {

			plist = append(plist[:i], plist[i+1:]...)
			return true
		}
	}

	return false
}

//uye fonk bitiş

//Profil ile alakalı fonk lar
func (u Uye) GetProfil(id bson.ObjectId) Profil {
	up := u.Get(id)
	return up.Profil

}

func (u Uye) UpdateProfil(id bson.ObjectId, p Profil) {
	up := u.Get(id)
	up.Profil = p
	u.Update(up)
}

//Profil Bitiş

//mesaj ile alakalı
func (m Mesaj) Create(msj Mesaj) {
	s := getSession()
	c := s.DB("test").C("mesaj")
	msj.Id = bson.NewObjectId()
	err := c.Insert(msj)
	if err != nil {
		log.Fatal(err)
	}
}

func (m Mesaj) Update(msj Mesaj) {
	s := getSession()
	c := s.DB("test").C("mesaj")
	err := c.Update(msj.Id, msj)
	if err != nil {
		log.Fatal(err)
	}
}
func (m Mesaj) Get(id bson.ObjectId) Mesaj {
	s := getSession()
	c := s.DB("test").C("mesaj")
	mg := Mesaj{}
	err := c.Find(bson.M{"_id": id}).One(&mg)
	//err := c.Find(bson.D{"_id": id}).One(&mg)
	if err != nil {
		log.Fatal(err)
	}
	return mg
}
func (m Mesaj) GetByUyeId(id bson.ObjectId) Mesaj {
	s := getSession()
	c := s.DB("test").C("mesaj")
	mg := Mesaj{}
	err := c.Find(bson.M{"konusanlar._id": id}).One(&mg)

	if err != nil {
		//log.Fatal(err)
	}
	return mg
}
func (m Mesaj) Insert(msj MesajInfo) {

	temp := m.Get(msj.Id)

	msj.Id = bson.NewObjectId()
	temp.Mesajlar[msj.Id.String()] = msj
	m.Update(temp)

}

//mesaj bitti

//dilller
func (us Dil) Create(u Dil) {
	u.Id = bson.NewObjectId()
	session := getSession()
	c := session.DB("test").C("diller")
	err := c.Insert(u)
	if err != nil {
		log.Fatal(err)
	}
}

func (us Dil) Get(id string) Dil {
	session := getSession()
	c := session.DB("test").C("diller")
	u := Dil{}
	err := c.Find(bson.M{"Id": bson.ObjectIdHex(id)}).One(&u)
	if err != nil {
		log.Fatal(err)
	}
	return u
}

func (us Dil) Update(u Dil) {
	session := getSession()
	c := session.DB("test").C("diller")
	err := c.Update(u.Id, u)
	if err != nil {
		log.Fatal(err)
	}
}
func (us Dil) Delete(d Dil) {

	//oid := bson.ObjectIdHex(id)
	session := getSession()
	c := session.DB("test").C("diller")
	if err := c.Remove(d); err != nil {
		return
	}
}

//dilller bitiş

///Ayarlar başlangç

func (u Uye) AyarlarGet(id bson.ObjectId) Ayar {
	a := Ayar{}
	tmp := u.Get(id)
	a = tmp.Ayarlar
	return a
}

func (u Uye) AyarlarUpdate(id bson.ObjectId, a Ayar) Ayar {
	tmp := u.Get(id)
	tmp.Ayarlar = a
	return a
}

///Ayarlar bitiş
///Arkadaşlık başlangıç
func (u Uye) ArkadaslikEkle(g_id bson.ObjectId, a_id bson.ObjectId) {
	a := Arkadaslik{}
	a.Alici = u.Get(a_id)
	a.Gonderen = u.Get(g_id)
	a.Durum = false
	a.Onay = false
	a.Tarih = time.Now()
	a.Id = bson.NewObjectId()

	gonderen := u.Get(g_id)
	gonderen.Istekler[a.Id.String()] = a
	u.Update(gonderen)

	alici := u.Get(g_id)
	alici.Istekler[a.Id.String()] = a
	u.Update(alici)

	session := getSession()
	c := session.DB("test").C("arkadaslik")
	if err := c.Insert(a); err != nil {
		return
	}

}

func (u Uye) ArkadaslikOnayla(id bson.ObjectId, istek_id bson.ObjectId) {

	gonderen := u.Get(id)
	a := gonderen.Istekler[istek_id.String()]

	alici := u.Get(a.Alici.Id)

	alici.Arkadaslar[gonderen.Id.String()] = gonderen
	u.Update(alici)

	a_g := u.Get(a.Gonderen.Id)
	a_g.Arkadaslar[alici.Id.String()] = alici

	u.Update(a_g)

}

func (u Uye) ArkadaslikSil(id bson.ObjectId) {
	session := getSession()
	c := session.DB("test").C("arkadaslik")
	if err := c.Remove(bson.M{"Id": id}); err != nil {
		return
	}
}

func (u Uye) ArkadaslikCikar(a_id bson.ObjectId, s_id bson.ObjectId) {

	uye := u.Get(a_id)
	delete(uye.Arkadaslar, s_id.String())
	u.Update(uye)

	uye = u.Get(s_id)
	delete(uye.Arkadaslar, a_id.String())
	u.Update(uye)

}

func (u Uye) ArkadaslikIstekleri(id bson.ObjectId) []Arkadaslik {
	var a []Arkadaslik
	session := getSession()
	c := session.DB("test").C("arkadaslik")
	if err := c.Find(bson.M{"Id": id}).All(&a); err != nil {
		log.Fatal(err)
	}
	return a
}

///Arkadaşlık bitiş
