# Repository

Design Pattern untuk memisahkan query logic (berhubungan langsung dengan database) dengan business logic.

## Kenapa Query Logic dan Business Logic harus dipisahkan?

Sekedar implementasi dari 'S' pada SOLID yaitu **Single Responsibility Principle** dimana 1 Class cuman boleh menjalankan satu task aja. Nah, jadi biar ga berceceran (kadang kalau query error kita gatau errornya ada dibagian mana). Kita satukan jadi 1 file repository per entity.
Misal, kita ada:
```md
user_entity.go
article_entity.go
.
.
.
x_entity.go
```
Nanti harus ada:
```md
user_repository.go
article_repository.go
.
.
.
x_entity.go
```
Biasanya juga kita kalau CRUD kan kek begini:
```php
class ArticleController {
  // constructor etc.
  public function getArticle() {
    $id = $_GET['id']
    $sql = "SELECT * FROM Article WHERE id = {$id}"

    // error handling...

    $row = mysql_fetch_row($result);

    // process row...
  }
}
```

## Cara Implementasi

Nah, repository pattern itu instead kita satuin pengambilan data ke database dengan logic lainnya di class yang sama, kita pisahkan jadi 2 class yang saling berhubungan.
```php
class ArticleController extends Controller {
  public function createArticle(Request $req){
    $id = $req->get('id');
    $repo = new ArticleRepository();
    $article = $repo->findById($id);
    return response()->json(['article'->...]);
  }
}

class ArticleRepository {
  public function findById($id){
    $query = "SELECT * FROM Article WHERE id = {$id}";
    $article = $query->getRow(0, \App\Entities\User::class);
    return $article;
  }
}
```
Nah, meskipun nambah ribet, at least kalau ada error di query, kita gaush pusing nyari baris querynya ada dimana. Ini juga kita ngelakuin yang namanya **Code Decoupling**.

## Manfaat

Apa sih manfaat dari Code Decoupling ini? Untuk mengetahuinya kita tinggal buat semacam contract pada class repository dimana contract ini bakal bertindak sebagai function-function apa aja yang harus ada dalam class Repository.
```php
type IArticleRepository interface {
  createArticle(...)
  ...
  findById($id)
}
```
Dengan begini ArticleController tidak perlu tau/peduli gimana logic dari setiap function yang ada, dia bisa langsung manggil functionnya bahkan jika function tersebut belum ada logicnya. Dengan begini juga, jika kita harus refactor method-method yang ada pada repository (misal pindah jenis database), class controller tidak akan terpengaruh selama contract dalam `IArticleRepository` tidak berubah.
