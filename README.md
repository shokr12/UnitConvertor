# Unit Converter (Go)

A **command-line Unit Converter** written in Go that allows you to easily convert between different units of **length, weight, and temperature**.  
This project demonstrates **basic Go programming, user input handling, reusable functions, and working with maps**.

---

## Features
- Convert **length units**: mm, cm, m, km, in, ft, yd, mi  
- Convert **weight units**: g, kg, lb, oz  
- Convert **temperature units**: Celsius (C), Fahrenheit (F), Kelvin (K)  
- Handles invalid input gracefully  
- Clear and interactive **CLI interface**  

---

## How to Run
1. Clone the repository:
```bash
git clone https://github.com/shokr12/UnitConvertor.git
Navigate to the project folder:

bash
Copy code
cd unit-converter
Run the application:

bash
Copy code
go run main.go
Usage
Run the program using go run main.go.

Choose the type of conversion:

1: Length

2: Weight

3: Temperature

Enter the value and units to convert from/to.

The program outputs the converted value.

Example
powershell
Copy code
Select conversion type:
1
Enter the value to convert:
5
Enter the unit to convert from (mm, cm, m, km, in, ft, yd, mi):
m
Enter the unit to convert to (mm, cm, m, km, in, ft, yd, mi):
ft
5.0000 m = 16.4042 ft
Challenges Solved
Created reusable functions for unit conversions.

Managed user input and errors effectively.

Built a clean CLI interface that’s easy to navigate.

Future Improvements
Add more unit types (volume, speed, area, etc.)

Add a graphical user interface (GUI) version

Add support for batch conversions via file input

Author
Mahmoud Shokr

GitHub: shokr12

Email: mshokr1dhdj@gmail.com

⚡ Fun fact: I love turning complex problems into simple, efficient solutions!
