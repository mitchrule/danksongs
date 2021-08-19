const JWT_TOKEN = localStorage.getItem('')

// var url:RequestInfo;
// var body:RequestInit;
var res:Response;
// var resJSON:JSON;

// Get  request that appends JWT token in required
export const get = async ( url: RequestInfo ) => {

    res = await fetch(url);
    const resJSON = await res.json();
    return resJSON;
}


// Post request that appends JWT token in required
export const post = async (url: RequestInfo, body: RequestInit) => {
    
    res = await fetch(url, body);
    const resJSON = await res.json();
    return resJSON;
}

