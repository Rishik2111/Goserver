Base URL : "http://localhost:9000"

Services Provided:
1. Add User
    - POST
    - Route : "/users/createuser"
    - Data {Name, Email, FavGenres}
2. Delete User
    - DELETE
    - Route : "/users/deleteuser/:User_Id"
3. Get User List
    - GET
    - Route : "/users/getall"
4. Get Particular User by Userid
    - GET
    - Route : "/users/getuser/:User_Id"


5. Add Movie
    - POST
    - Route : "/movies/createmovie"
    - Data {Title, Genre, Storyline}
6. Delete Movie
    - DELETE
    - Route : "/movies/deletemovie/:Movie_Id"
7. Get Movie List
    - GET
    - Route : "/movies/getall"
8. Get Particular Movie by Movieid
    - GET
    - Route : "/movies/getmovie/:Movie_id"


9. Add a movie to user's movie list
    -POST
    -Route : "/usermovie/add"
    -Data { UserId, MovieId, Rating(Optional)}

10. Delete a movie from user's movie list by providing Userid, Movieid
    -DELETE
    -Route : "/usermovie/deleteusermovie"
    -Params : {userid, movieid}

11. Get the movie list of a User by Userid
    - GET
    - Route : "/usermovie/getusermovies/:User_id"

