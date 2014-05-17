package controllers
import "strconv"
import "github.com/menghx/mxgo"

type ErrorController struct {
	*mxgo.Controller
}

func (errc *ErrorController)Handle(errorCode string) Result{
	ec,err := strconv.Atoi(errorCode)
	if err {
		panic(err)
	}
	switch {
		case ec>=300&&ec<400:{
			return handle4xx(ec)
		}
		case ec>=400&&ec<500:{
			return handle4xx(ec)
		}
		case ec>=500&&ec<600:{
			return handle5xx(ec)
		}
	}
	return handle(ec)
}

func handle(ecode int) Result{
	return nil
}

func handle3xx(ecode int) Result{
	return nil
}

func handle4xx(ecode int) Result{
	return nil
}

func handle5xx(ecode int) Result{
	return nil
}