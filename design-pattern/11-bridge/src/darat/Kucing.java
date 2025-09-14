package darat;

import bridges.BinatangDarat;

public class Kucing extends BinatangDarat {
    @Override
    public int getJumlahKaki() {
        return 4;
    }

    @Override
    public String getNama() {
        return "Kucing";
    }
}
