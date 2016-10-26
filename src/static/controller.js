var app = angular.module("TweetsDelete", []);

app.controller("TweetsDeleteController", function ($scope, $http) {
    $scope.extensions = ["CSV", "JSON"];
    $scope.favoritedChoices = ["No", "Yes"];

    $scope.canSend = function () {
        return !($scope.consumerKey && $scope.consumerSecret
        && $scope.accessToken && $scope.accessSecret
        && $scope.userName);
    };

    $scope.deleteAllTweets = function () {
        console.log($scope.keepFavorited);

        var deleteUrl = "deleteAll";

        if ($scope.keepFavorited === "Yes") {
            deleteUrl = "deleteAllExceptFavorited";
        }

        $http({
            method: "POST",
            url: "http://localhost:8080/twitterDelete/" + deleteUrl,
            data: {
                "consumerKey": $scope.consumerKey,
                "consumerSecret": $scope.consumerSecret,
                "accessToken": $scope.accessToken,
                "accessSecret": $scope.accessSecret
            }
        }).success(function (responseData) {
            $scope.deleteMessage = responseData;
        }).error(function (data) {
            $scope.result = "";
        });
    };

    $scope.exportAllTweets = function () {
        $http({
            method: "POST",
            url: "http://localhost:8080/twitterDelete/export",
            data: {
                "userName": $scope.userName,
                "extension": $scope.extension,
                "credentials": {
                    "consumerKey": $scope.consumerKey,
                    "consumerSecret": $scope.consumerSecret,
                    "accessToken": $scope.accessToken,
                    "accessSecret": $scope.accessSecret
                }
            }
        }).success(function (responseData) {
            $scope.exportFilePath = responseData;
        }).error(function (data) {
            $scope.result = "";
        });
    };
});
