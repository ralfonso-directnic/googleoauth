package main

import (
  "html/template"
  ga "github.com/ralfonso-directnic/googleoauth"
  "os"
  "github.com/markbates/goth"
  "net/http"
  "log"
)


var loginTemplate = `<a href="/auth/google">Sign in</a>`



var userTemplate = `
<p><a href="/logout/{{.Provider}}">logout</a></p>
<p>Name: {{.Name}} [{{.LastName}}, {{.FirstName}}]</p>
<p>Email: {{.Email}}</p>
<p>NickName: {{.NickName}}</p>
<p>Location: {{.Location}}</p>
<p>AvatarURL: {{.AvatarURL}} <img src="{{.AvatarURL}}"></p>
<p>Description: {{.Description}}</p>
<p>UserID: {{.UserID}}</p>
<p>AccessToken: {{.AccessToken}}</p>
<p>ExpiresAt: {{.ExpiresAt}}</p>
<p>RefreshToken: {{.RefreshToken}}</p>`


func main(){
    
    os.Setenv("SESSION_SECRET","123445646")
    os.Setenv("GOOGLEOAUTH_HOST","mymachinehost.com")
    os.Setenv("GOOGLEOAUTH_KEY","madeupstringhere.apps.googleusercontent.com")
    os.Setenv("GOOGLEOAUTH_SECRET","abc123")
    os.Setenv("GOOGLEOAUTH_PORT","8080")
    
    ga.Config()
    ga.AuthListen(loginTemplate,func(user goth.User,res http.ResponseWriter,req *http.Request){
        
         log.Printf("%+v",user)
         t, _ := template.New("foo").Parse(userTemplate)
		 t.Execute(res,user)
		 
		 //write out a result, redirect,save a session,etc
        
    })
    
    
}
