import bridges.BinatangDarat;
import bridges.BinatangLaut;
import darat.Anjing;
import darat.Kucing;
import darat.Monyet;
import laut.Hiu;
import laut.Koi;
import laut.Lele;
import parent.Binatang;

import java.util.ArrayList;

public class Main {
    public static void main(String[] args) {
        ArrayList<Binatang> list = new ArrayList<>();
        list.add(new Anjing());
        list.add(new Kucing());
        list.add(new Monyet());
        list.add(new Hiu());
        list.add(new Koi());
        list.add(new Lele());

        for(Binatang binatang : list) {
            if (binatang instanceof BinatangDarat) {
                BinatangDarat binatangDarat = (BinatangDarat) binatang;
                System.out.println(binatangDarat.getNama() + " hidup di darat dengan " + binatangDarat.getJumlahKaki() + " kaki.");
            } else if (binatang instanceof BinatangLaut) {
                BinatangLaut binatangLaut = (BinatangLaut) binatang;
                System.out.println(binatangLaut.getNama() + " hidup di laut.");
            }
        }
    }
}