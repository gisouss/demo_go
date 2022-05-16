GO RESful API demo


<ins>Prerequisite</ins>:

 - GO must be installed <link>https://go.dev/dl/</link>
 
<ins>Execution</ins>: 

 - Get dependencies:
    <code>go get .</code>
 
 - Run Gin web server:
    <code>go run .</code>
 
 <ins>BugFix</ins>:
 
 If you are running the demon on a Mac, you have to execute:
 <code>go get -u golang.org/x/sys</code>

<ins>Initial data</ins>:

<code>{ID: "1", Tech: "Python", Rate: "5/5"}</code>
<code>{ID: "2", Tech: "JavaScript", Rate: "4/5"}</code>
<code>{ID: "3", Tech: "C#", Rate: "3/5"}</code>


<ins>Routes examples</ins>:

 - Get all technologies : <link>http://localhost:8081/technologies</link>
 - Get technology by id (id=1) <link>http://localhost:8081/technologies/1</link>
 - Add technology <link>http://localhost:8081/technologies</link>
 Payload:
 <code>{
    "id": "3",
    "tech": "C#",
    "rate": "3/5"
  }</code>
- Update technology by id (id=2) <link>http://localhost:8081/technologies/2</link>
 Payload:
 <code>{
    "id": "2",
    "tech": "Angular",
    "rate": "2/5"
  }</code>
 - Delete tehnology by id (id=1) <link>http://localhost:8081/technologies/1</link>
