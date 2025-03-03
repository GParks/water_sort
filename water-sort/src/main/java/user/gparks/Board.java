package user.gparks;

class Move {
    int src;
    int dst;
    ColorEnum color;
    public Move(int src, int dst, ColorEnum color) {
        this.src = src;
        this.dst = dst;
        this.color = color;
    }
}


public class Board implements Comparable<Board> {
    protected final int NUM_TUBES = 14;

    @Override
    public int compareTo(Board o) {
        return 0;
    }

    
}
