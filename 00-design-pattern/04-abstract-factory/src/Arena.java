interface Arena {
    public void start();
}

class ArenaEasy implements Arena{
    @Override
    public void start() {
        System.out.println("Arena Easy");
    }
}

class ArenaMedium implements Arena{
    @Override
    public void start() {
        System.out.println("Arena Medium");
    }
}

class ArenaHard implements Arena{
    @Override
    public void start() {
        System.out.println("Arena Hard");
    }
}
