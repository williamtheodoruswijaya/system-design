public class BlockGame extends BlockGameTemplateMethod{
    @Override
    public String getTitle() {
        return "BLOCK GAME 1";
    }

    @Override
    public String getEndTitle() {
        return "FINISH PLAYING BLOCK GAME";
    }

    @Override
    public char getCharacter() {
        return '*';
    }

    @Override
    public int getHeight() {
        return 10;
    }

    @Override
    public int getWidth() {
        return 10;
    }
}
