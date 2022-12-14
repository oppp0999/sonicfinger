package routers

import "fmt"

func game(){
	prepare();
	// 스페이스바를 눌러 게임을 시작할 수 있도록 키보드 이벤트 등록
	fmt.Sprintf("%s", )
	$(document.body).keyup(start);
	// 입력된 값으로 단어 카드를 제거하기 위한 이벤트 등록
	$("#target").keyup(doClear);
}


// 단어 카드가 화면에 출력되는 인터벌을 밀리세컨드로 설정.
// 이 시간을 조정함으로써 게임의 레벨을 정할 수 있다.
var intervalTime = 1500;
// 최대 실패 허용 단어 카드 수 설정.
var limitedCount = 3;
// 준비된 단어를 정해진 시간에 반복적으로 화면에 출력하는 처리를 위해 사용되는 setInterval() 메소드의 프로세스 ID를 대입하기 위한 변수
var makeProc;
// 반복적으로 성공 여부를 검증하기 위해 사용하는 setInterval() 메소드의 프로세스 ID를 대입하기 위한 변수
var completedProc;
// 활성 단어 카드 갯수 저장을 위한 변수 준비.
var count;
// 단어 카드를 배열에 담기 위한 배열 변수 준비
var targetList;
// 출력된 단어 카드를 저장하기 위한 변수 준비(중복된 단어 카드 출력 방지).
var applyList = [];
// 커버 요소를 jQuery 선택자로 저장하기 위한 변수 준비.
var $cover;
// http://ko.talkenglish.com/vocabulary/top-2000-vocabulary.aspx
var words = "the,of,and,to,in,you,for,or,it,as,be,on,with,can,have,this,by,not,but,at,from,they,more,will,if,some,there,what,about,which,when,one,all,also";
/**
 * 게임을 처음 시작할 때 준비를 하는 함수.
 */
func prepare() {
$('<div id="cover">').appendTo(document.body);
$cover = $("#cover");
var html = '<h1>타자 연습 게임</h1>';
html += '<p>시작하면 단어 카드가 제시되며 단어를 그대로 입력후 엔터를 누르시면 제거됩니다.<br />';
html += limitedCount + '개의 단어 카드를 제거하지 못하고 남게 되면 게임은 종료됩니다. .</p>';
html += '<p>지금 게임을 시작하려면 스페이스바를 누르세요.</p>';
$cover.addClass("cover").html(html);
}

/**
 * 게임 시작을 하는 함수
 * 게임에 필요한 데이터를 재설정한다.
 */
function start(event) {
var keyCode = event.which || event.keyCode;
// 스페이스바를 누른 경우.
if (keyCode == 32) {
// 혹시 모를 남아 있는 단어 카드를 모두 제거한다.
$(".game-target-container").children().remove();
// 단어 카드 배열을 재설정한다.
//targetList = ['apple', 'banana', 'kwie', 'melon', 'mango', 'friday'];
targetList = words.split(",");
// 출력된 단어 카드 목록 초기화.
applyList = [];
// 활성 단어 카드수를 초기화.
count = 0;
// 커버를 닫느다.
$("#cover").hide();

// 단어 카드를 만드는 반복 처리 프로세스 등록.
makeProc = setInterval(makeTarget, intervalTime);
// 1/2초마다 게임의 성공 여부를 체크하도록 반복 처리 프로세스 등록
completedProc = setInterval(isCompleted, 500);
// 단어 카드를 지우기 위한 입력 필드를 사용할 수 있도록 준비.
$("#target").val("").attr("disabled", false).focus();
}
}

/**
 * 단어 카드를 만드는 함수.
 */
function makeTarget() {
// 먼저 제거되지 않은 단어 카드의 갯수를 체크해 미션 실패 여부를 확인.
isCompleted();
// 준비된 단어의 갯수를 저장.
var totalItem = targetList.length;
// 더 이상 단어 카드를 만들 자료가 없는 경우에 대한 처리.
// count는 기본적으로 생성후 1이 더해지므로 큰 경우로 조건 검사를 한다.
if (count > totalItem - 1) {
clearInterval(makeProc);
return;
}

// 랜덤하게 인덱스 생성.
var idx = getRandomInt(0, totalItem - 1);

// 단어카드를 만들 단어 저장.
var item = targetList[idx];
if (applyList.indexOf(item) < 0) {
var target = '<span class="' + item + '">' + item + '</span>';
$(".game-target-container").append(target);
applyList.push(item);
// 활성 단어 카드수 1 증가.
count++;
}

}

/**
 * 단어 카드 제거 함수.
 */
function doClear(event) {
// 이벤트 버블링을 중지.
event.stopPropagation();

var keyCode = event.which || event.keyCode;
// enter키를 누른 경우.
if (keyCode == 13) {
var target = $(this).val();
var $target = $("." + target);

if ($target.length) {
$target.remove();
var idx = targetList.indexOf(target);
if (idx > -1) {
// 배열 제거.
targetList.splice(idx, 1);
// 활성 단어 카드수 1 감소.
count--;
}
}

// 값 비움.
$(this).val("");

// 성공 어부를 체크
isCompleted();
}
}

/**
 * 미션 성공 여부를 체크
 */
function isCompleted(result) {
var result;
// 남아 있는 단어 카드의 갯수가 실패 허용 갯수를 넘은 경우에 실패 처리
if (count > limitedCount - 1) {
result = "failure";
// 단어 카드를 만들 단어의 갯수가 0이거나 현재 남아있는 단어의 갯수가 최대 실패 허용 갯수보다 작은 경우에 미션 성공 처리.
} else if (targetList.length == 0 || targetList.length < limitedCount) {
result = "completed";
} else {
result = null;
}

if (result) {
$("#target").attr("disabled", true);
clearInterval(makeProc);
clearInterval(completedProc);

var html;
if (result == "failure") {
html = '<h2>실패했습니다.</h2>';
} else {
html = '<h2>성공했습니다.</h2>';
}
html += '<p>게임을 다시 시작하려면 스페이스바를 누르세요.</p>';
$cover.html(html).show();

return;
}
}

/**
 * 랜덤 숫자 구하는 함수.
 */
function getRandomInt(min, max) {
return Math.floor(Math.random() * (max - min + 1)) + min;
}
