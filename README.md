# ASCII-ART WEB

Ascii-art-web consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) version of the last project ascii-art.


## Description
Ascii-art Web is a program that creates a webpage for clients to generate Ascii-Art of the message you enter. There are various aspects that can be implemented where the user/client is presented an area for them to type their desire message. If the character is among the English Dictionary it is represented within the art as well as numerics and symbols.






## Usage
* You have to install go dependancies for you to run the program. If you don't have the go dependancies run the following command to download ;

`sudo apt install go -all`

* If you have go dependancies skip the first step. To clone the project on your local machine, run

`git clone https://learn.zone01kisumu.ke/git/shfana/ascii-art-web.git`

* Navigate to the directory where the project is and run the program

```bash
cd ascii-art-web
go run .
```
output

You will receive the following message on your terminal
```bash
Server is starting at http://localhost:8080
```
To view the project on your browser, you can click the link on the terminal or copy it in your browser, or open localhost at port **8080**

This will be the home page of our application when opened on the browser(Mozilla Firefox)

![home](/assets/images/home.png)

- This is the page that is displayed when the user opens [http://localhost:8080](http://localhost:8080)

- These are the various parts of the home page 

![instructions](/assets/images/instruct.png)


 + Navbar - The navbar links to other pages of the application

 + The Input section - where the user keys in the input they want to generate to art

 + banner selection - The user has three options of banner files to choose from. (shadow, thinkertoy and standard)

 + The submit button - The user presses this button to generate the art that is displayed in the output section.

 + The footer  - This section contains the copyright disclaimer and an icon that has a  link to the repository.


![banner](/assets/images/banner.png)


The user can select one of the three available fonts/banner files to display their input in art form.

This project comes with three predefined ASCII fonts, located in banner files:

  - [`shadow`](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/shadow.txt)
  - [`standard`](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/standard.txt)
  - [`thinkertoy`](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/thinkertoy.txt)


**- Each character is represented over 8 lines, providing a clear and sizable output.**

**- Characters are separated by a new line `\n`.**

- From the example where the user has typed "hello" in the input field and selected  the `standard` banner. 
- The output should be shown as seen in the image below
![output](/assets/images/outputS.png)

- In the case of input `Hello` and `shadow` banner selected, the output is as shown below
![shadow](/assets/images/ouputSh.png)


- In the case of input `Hello` and `thinkertoy` banner selected.
![thinkertoyinput](/assets/images/thinkertoy%20input.png)

- Output is as shown below
![thnkertoyoutput](/assets/images/thinkertou%20output.png)

- Bad Requests: They occur when

    - The user does not give an input
    - The user gives an input that is not within the printable ASCII character range (from space to tide `~`). There are 95 characters in total

An example of a bad request is shown below(the user tries to input chinese characters or emojis)
![badrequest](/assets/images/badrequest.png)

The Following error is returned.
![badreq](/assets/images/badrequestout.png)

The user can click `go back` to return home to try using the program again


- In the case where a user tries to navigate to a non existent page or use a banner file that is not  available, he/she gets a 404 not found error 
![notfound](/assets/images/not%20found.png)

- Internal server error occurs when an error occurs from the program and is not a user error.
![internalservererror](/assets/images/ISE.png)

## Implementation detail: algorithm
* ### assets




* ### functions



* ### server




* ### template


* ### main.go



## Limitations
This project can only print characters in the printable ASCII range from space to tide `~`. When other characters outside the range are used, the program returns a bad request error.

## Contribution

Contributions are welcome. Please adhere to the existing coding standards and include unit tests for any new features or changes. Ensure to thoroughly test the code before pushing any updates.
If you encounter any issues or have suggestions for improvement, feel free to submit an issue, pull request or propose a change!


## License
This Program is under the **MIT** Licence see [LICENSE](/LICENSE) for details.

## Authors
This program was built and maintained by 
- [cliffootieno](https://learn.zone01kisumu.ke/git/cliffootieno)
- [wnjuguna](https://learn.zone01kisumu.ke/git/wnjuguna)
- [shfana](https://learn.zone01kisumu.ke/git/shfana)