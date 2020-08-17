package googleoauth

import (
	"fmt"
	"html/template"
	"net/http"
	"log"
	"github.com/gorilla/pat"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
        "io/ioutil"
        "os"
 	"github.com/markbates/goth/providers/google"
)





func Config() {
  
    if(os.Getenv("GOOGLEOAUTH_HOST")==""){
      hst,_ := os.Hostname()
      os.Setenv("GOOGLEOAUTH_HOST",hst)  
    }
    
   if(os.Getenv("GOOGLEOAUTH_PORT")==""){

      os.Setenv("GOOGLEOAUTH_PORT","3000")  
    }
    
    
    

	goth.UseProviders(
      

		google.New(os.Getenv("GOOGLEOAUTH_KEY"), os.Getenv("GOOGLEOAUTH_SECRET"), "http://"+os.Getenv("GOOGLEOAUTH_HOST")+":"+os.Getenv("GOOGLEOAUTH_PORT")+"/auth/google/callback"),
	)
	

}

func AuthListen(loginTemplate string,fn func(user goth.User,res http.ResponseWriter,req *http.Request)){
    
    
    //could be a path or could be a string
    
    f, err := ioutil.ReadFile(loginTemplate)

    if err == nil {
       loginTemplate = string(f) // we found a file with path loginTemplate so we set this to the new string, otherwise it's just a string template
    }

    

	p := pat.New()
	p.Get("/auth/{provider}/callback", func(res http.ResponseWriter, req *http.Request) {

		user, err := gothic.CompleteUserAuth(res, req)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		
    	
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
	
         	    fn(user,res,req)
    
	})

	p.Get("/logout/{provider}", func(res http.ResponseWriter, req *http.Request) {
		gothic.Logout(res, req)
		res.Header().Set("Location", "/")
		res.WriteHeader(http.StatusTemporaryRedirect)
	})

	p.Get("/auth/{provider}", func(res http.ResponseWriter, req *http.Request) {

		
		if gothUser, err := gothic.CompleteUserAuth(res, req); err == nil {
		
		     	    fn(gothUser,res,req)
		
		} else {
			gothic.BeginAuthHandler(res, req)
		}
	})

	p.Get("/", func(res http.ResponseWriter, req *http.Request) {
    	
    	
		t, _ := template.New("foo").Parse(loginTemplate)
		t.Execute(res,nil)
	})

	//log.Println("listening on "+os.Getenv("GOOGLEOAUTH_HOST")+":"+os.Getenv("GOOGLEOAUTH_PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("GOOGLEOAUTH_PORT"), p))
}




