public class ScreencastAdapter implements  CatalogAdapter {
    private Screencast screencast;

    public ScreencastAdapter(Screencast screencast) {
        this.screencast = screencast;
    }

    @Override
    public String getCatalogTitle() {
        return screencast.getTitle() + " by " + screencast.getAuthor() + " -> " + screencast.getDuration();
    }
}
