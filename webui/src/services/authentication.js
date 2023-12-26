function getCurrentId(){
    return localStorage["token"]
}

function getCurrentUsername(){
    return localStorage["username"]
}

export {
    getCurrentId, 
    getCurrentUsername,
}