public abstract class BlockGameTemplateMethod {
    public void start() {
        System.out.println(getTitle());

        for(int i = 0; i < getHeight(); i++) {
            for(int j = 0; j < getWidth(); j++) {
                System.out.print(getCharacter());
            }
            System.out.println();
        }

        System.out.println(getEndTitle());
    }

    public abstract String getTitle();

    public abstract String getEndTitle();

    public abstract char getCharacter();

    public abstract int getHeight();

    public abstract int getWidth();
}
