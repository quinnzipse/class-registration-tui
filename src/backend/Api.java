import java.util.*;

public class Api{
  public static void main(String[] args){
    Scanner s = new Scanner(System.in);

    List<Course> courseList = new ArrayList<Course>();
    // List<Professor> professorList = new ArrayList<Professor>();

    while(s.hasNextLine()){
      String line = s.nextLine();
      // System.out.println(line);
      String[] tokens = line.split(" ");

      switch(tokens[0]){
        case "get-classes":
          StringBuilder sb = new StringBuilder();
          courseList.forEach(c -> sb.append(c.toString()).append(" "));
          System.out.println("ok. " + sb.toString());
          break;

        // add-class Professor ClassName End-Time Start-Time DaysOfWeek        
        case "add-class":

          if(tokens.length < 6){
            System.out.println("lol no.");
            break;
          }

          int dsow = 0;
          
          if(tokens[5].contains("Mo")) dsow |= 1;
          if(tokens[5].contains("Tu")) dsow |= 2;
          if(tokens[5].contains("We")) dsow |= 4;
          if(tokens[5].contains("Th")) dsow |= 8;
          if(tokens[5].contains("Fr")) dsow |= 16;

          Course c = new Course(tokens[1], tokens[2], dsow, tokens[3], tokens[4]);
          courseList.add(c);
          System.out.println("ok.");
          break;
        case "delete-class":
          if(tokens.length < 2){
            System.out.println("lol no.");
            break;
          }

          Course remove_this_guy = null;
          for(int i=0; i<courseList.size(); i++){
              if(courseList.get(i).Cs_class.equals(tokens[1])){
                remove_this_guy = courseList.get(i);
                break;
              }
          }

          boolean removed = courseList.remove(remove_this_guy);

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
