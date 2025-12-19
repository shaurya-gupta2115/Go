package auth


func extractSession() string{ // yaha pr is method ko start isliye aisa kiya hai kyunki is hum PRIVATE RAKHNA HAI
	return "loggedIn"
}

func GetSession() string{ // ye PUBLIC rakhna hai isiliye iska pehla word capital rakha hai 
	return extractSession()
}

