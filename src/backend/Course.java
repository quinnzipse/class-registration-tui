//hard example 

public class Course {
	//private global variables 
	String ProfessorName;
	String Cs_class;
	int days;
	String endTime;
	String startTime;
	
	//Constructor
	//two string method
	
	
	public Course(String ProfessorName2, String Cs_class, int days, String endTime2, String startTime2) {
		this.Cs_class =  Cs_class;
		this.startTime = startTime2;
		this.endTime = endTime2;
		this.days = days; 
		this.ProfessorName = ProfessorName2; 
		
	}
	
	public String toString() {
		return ProfessorName + " " + Cs_class + " " + days + " " + endTime + " " + startTime; 
	}
}

	