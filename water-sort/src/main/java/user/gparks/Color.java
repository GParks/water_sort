package user.gparks;
import java.util.HashMap;

/** 
 * Enumeration of Colors
 */
enum ColorEnum {
    NC, 
    RED, 
    GREEN, OLIVE, 
    LTBLUE, BLUE, 
    YELLOW, 
    PURPLE, 
    ORANGE, 
    BROWN, PINK, 
    NEONGREEN, 
    GRAY;
    // https://stackoverflow.com/questions/609860#19277247
    public static ColorEnum values[] = values();
}

public class Color {
    protected String[] colorNames = {
        "*** No Color ***", 
        "Red", 
        "Green", "Olive", 
        "LiteBlue", "Blue", 
        "Yellow", 
        "Purple", 
        "Orange", 
        "Brown", "Pink", 
        "NeonGreen", 
        "Gray"
    };

    protected HashMap<String, ColorEnum > colorMap; 
    public Color() {
        colorMap = new HashMap<String, ColorEnum>();
        for (int i = 0; i < colorNames.length; i++) {
            colorMap.put(colorNames[i], ColorEnum.values[i]);
        }
    }

    public ColorEnum getColor(String color) {
        ColorEnum retval = colorMap.get(color);
        if (retval == null) {
            retval = ColorEnum.NC;
        }   
        return retval;
    }
    public String colorName(ColorEnum ce) {
        String retval = "NC";
        retval = colorNames[ce.ordinal()];
        return retval;
    }

}
