package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	// User represents the structure of our resource
	User struct {
		Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Isım     string        `json:"adsoyad" bson:"adsoyad"`
		Email    string        `json:"email" bson:"email"`
		Sifre    string        `json:"sifre" bson:"sifre"`
		Dtarihi  string        `json:"dtarihi" bson:"dtarihi"`
		Cinsiyet string        `json:"cinsiyet" bson:"cinsiyet"`
		Dil      string        `json:"anadil" bson:"anadil"`
		Ulke     string        `json:"ulke" bson:"ulke"`
		Il       string        `json:"il" bson:"il"`
		Presim   string        `json:"presim" bson:"presim"`
	}

	Profil struct {
		Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
		AdSoyad  string        `json:"adsoyad" bson:"adsoyad"`
		Yas      int           `json:"yas" bson:"yas"`
		Meslek   string        `json:"meslek" bson:"meslek"`
		Aciklama string        `json:"aciklama" bson:"aciklama"`
		Resimurl string        `json:"resimurl" bson:"resimurl"`
		//diller   []DilInfo			`json:"diller" bson:"diller"`
		Anadil      Dil               `json:"anadil" bson:"anadil"`
		Ogr_diller  map[string]Dil    `json:"odiller" bson:"odiller"`
		Biln_diller map[string]Dil    `json:"bdiller" bson:"bdiller"`
		Ulke        string            `json:"ulke" bson:"ulke"`
		Sehir       string            `json:"sehir" bson:"sehir"`
		Hobiler     string            `json:"hobiler" bson:"hobiler"`
		Arkadaslar  map[string]Profil `json:"arkadaslar" bson:"arkadaslar"`
		Hakkinda    string            `json:"hakkinda" bson:"hakkinda"`
		Amac        string            `json:"amac" bson:"amac"`
		Cinsiyet    string            `json:"cinsiyet" bson:"cinsiyet"`
		Dogumtarihi string            `json:"dtarihi" bson:"dtarihi"`
	}

	Uye struct {
		Id         bson.ObjectId         `json:"id" bson:"_id,omitempty"`
		Isım       string                `json:"adsoyad" bson:"adsoyad"`
		Email      string                `json:"email" bson:"email"`
		Sifre      string                `json:"sifre" bson:"sifre"`
		Profil     Profil                `json:"profil" bson:"profil"`
		Arkadaslar map[string]Uye        `json:"arkadaslar" bson:"arkadaslar"`
		Mesajlar   map[string]Mesaj      `json:"mesajlar" bson:"mesajlar"`
		Dil        Dil                   `json:"dil" bson:"dil"`
		Durum      bool                  `json:"durum" bson:"durum"`
		Ayarlar    Ayar                  `json:"ayar" bson:"ayar"`
		Aktif      bool                  `json:"aktif" bson:"aktif"`
		Istekler   map[string]Arkadaslik `json:"istekler" bson:"istekler"`
	}
	Ayar struct {
		Id            bson.ObjectId `json:"id" bson:"_id,omitempty"`
		ProfilGorunum bool          `json:"pgorunum" bson:"pgorunum"`
		MesajGonderim int           `json:"mgonderim" bson:"mgonderim"`
	}
	DilInfo struct {
		Id     bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Dil    Dil           `json:"dil" bson:"dil"`
		Seviye int           `json:"seviye" bson:"seviye"`
	}
	Dil struct {
		Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Ad       string        `json:"ad" bson:"ad"`
		Resimurl string        `json:"resimurl" bson:"resimurl"`
	}
	Mesaj struct {
		Id         bson.ObjectId        `json:"id" bson:"_id,omitempty"`
		Konusanlar map[string]Uye       `json:"konusanlar" bson:"konusanlar"`
		Mesajlar   map[string]MesajInfo `json:"mesajlar" bson:"mesajlar"`
		Ozet       string               `json:"ozet" bson:"ozet"`
		Tarih      time.Time            `json:"tarih" bson:"tarih"`
	}

	MesajInfo struct {
		Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Gonderen Uye           `json:"gonderen" bson:"gonderen"`
		Mesaj    string        `json:"mesaj" bson:"mesaj"`
		Tarih    time.Time     `json:"tarih" bson:"tarih"`
	}
	Sorgu struct {
		//Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Isim     string `json:"isim" bson:"isim"`
		Cinsiyet string `json:"cinsiyet" bson:"cinsiyet"`
		AzYas    int    `json:"azyas" bson:"azyas"`
		CokYas   int    `json:"cokyas" bson:"cokyas"`
		Online   bool   `json:"online" bson:"online"`
		Ulke     string `json:"ulke" bson:"ulke"`
		Dil      string `json:"dil" bson:"dil"`
	}

	Arkadaslik struct {
		Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Gonderen Uye           `json:"gonderen" bson:"gonderen"`
		Alici    Uye           `json:"alici" bson:"alici"`
		Tarih    time.Time     `json:"tarih" bson:"tarih"`
		Durum    bool          `json:"durum" bson:"durum"`
		Onay     bool          `json:"onay" bson:"onay"`
	}
)
