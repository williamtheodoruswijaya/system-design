public class BookAdapter implements CatalogAdapter{
    private Book book;

    public BookAdapter(Book book) {
        this.book = book;
    }

    @Override
    public String getCatalogTitle() {
        return book.getTitle() + " by " + book.getAuthor();
    }
}
