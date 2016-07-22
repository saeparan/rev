// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}


type tBoard struct {}
var Board tBoard


func (_ tBoard) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Board.Index", args).Url
}

func (_ tBoard) List(
		paging interface{},
		article interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "paging", paging)
	revel.Unbind(args, "article", article)
	return revel.MainRouter.Reverse("Board.List", args).Url
}

func (_ tBoard) View(
		article interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "article", article)
	return revel.MainRouter.Reverse("Board.View", args).Url
}

func (_ tBoard) Write(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Board.Write", args).Url
}

func (_ tBoard) Regist(
		article interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "article", article)
	return revel.MainRouter.Reverse("Board.Regist", args).Url
}


type tOAuth struct {}
var OAuth tOAuth


func (_ tOAuth) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("OAuth.Index", args).Url
}

func (_ tOAuth) Auth(
		code string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "code", code)
	return revel.MainRouter.Reverse("OAuth.Auth", args).Url
}


type tUser struct {}
var User tUser


func (_ tUser) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Index", args).Url
}

func (_ tUser) LoginSign(
		user interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("User.LoginSign", args).Url
}

func (_ tUser) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Logout", args).Url
}

func (_ tUser) Regist(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Regist", args).Url
}

func (_ tUser) Login(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("User.Login", args).Url
}

func (_ tUser) Add(
		user interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("User.Add", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


