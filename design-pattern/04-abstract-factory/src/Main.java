//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.
public class Main {
    public static void main(String[] args) {
//        Cara yang salah
//        Game gameEasy = new Game(new LevelEasy(), new ArenaEasy());
//        gameEasy.start();
//
//        Game gameMedium = new Game(new LevelMedium(), new ArenaMedium());
//        gameMedium.start();
//
//        Game gameHard = new Game(new LevelHard(), new ArenaMedium());
//        gameHard.start();

        Game gameEasy = new Game(new GameFactoryEasy());
        gameEasy.start();

        Game gameMedium = new Game(new GameFactoryMedium());
        gameMedium.start();

        Game gameHard = new Game(new GameFactoryHard());
        gameHard.start();
    }
}