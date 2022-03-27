import java.util.*;

public class Api{
  public static void main(String[] args){
    Scanner s = new Scanner(System.in);

    List<Course> courseList = new ArrayList<Course>();
    List<Professor> professorList = new ArrayList<Professor>();

    while(s.hasNextLine()){
      String line = s.getNextLine();
      String[] tokens = line.split(" ");

      switch(tokens[0]){
        case "get-classes":
          StringBuilder sb = new StringBuilder();
          courseList.forEach(c -> sb.append(c.toString()).append(" "));
          System.out.println("ok. " + sb.toString());
          break;
        
        case "add-class":
          // name, professor, endtime, starttime, days
          int[] dsow = new int[5];
          
          if(tokens[5].contains("Mo")) dsow[0] = 1;
          if(tokens[5].contains("Tu")) dsow[1] = 1;
          if(tokens[5].contains("We")) dsow[2] = 1;
          if(tokens[5].contains("Th")) dsow[3] = 1;
          if(tokens[5].contains("Fr")) dsow[4] = 1;

          Course c = new Course(tokens[1], tokens[2], tokens[3], tokens[4], dsow);
          courseList.add(c);
          System.out.println("ok.");
          break;
        case "delete-class":
          boolean removed = courseList.remove(tokens[1]);

          if(removed) System.out.println("ok.");
          else System.out.println("lol no.");
          break;
        case "get-professors":
          break;
        case "add-professor":
          break;
        case "delete-professor":
          break;
        default:
          System.out.println("Command Not Recognized!");
      }
    }  
  }
}
