const {SearchGameRequest, SearchGameResponse, GameSearch} = require('./generated/service/game/gamepb/game_pb.js');
const {GameServiceClient} = require('./generated/service/game/gamepb/game_grpc_web_pb.js');

var gameService = new GameServiceClient('http://localhost:50053');


var isSearching = false;
searchGame = function (term) {
    var term = document.getElementById('term').value;
    if (!term) {
        return;
    }

    document.getElementById('searchBtn').disabled = true;
    document.getElementById('searchList').innerHTML = '';

    var gameSearch = new GameSearch();
    gameSearch.setSearch(term);

    var request = new SearchGameRequest();
    request.setSearch(gameSearch);

    var list = document.getElementById('searchList');
    gameService.searchGame(request, {}, function (err, response) {
        if (err) {

        } else {
            var games = response.getGamesList();
            for (game of games) {
                var li = document.createElement("li");
                li.appendChild(document.createTextNode(game.getTitle()));
                list.appendChild(li);
            }
        }
        document.getElementById('searchBtn').disabled = false;
    });

};