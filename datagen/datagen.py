import argparse
import io
import json
import math
import random
import sys

import faker


parser = argparse.ArgumentParser(prog="datagen")
parser.add_argument(
    "-s",
    dest="nspots",
    type=int,
    default=1,
    help="number of spots",
)
parser.add_argument(
    "--out-spots",
    dest="out_spots",
    type=argparse.FileType("w"),
    help="output orders file name",
)
parser.add_argument(
    "-i",
    dest="nitems",
    type=int,
    default=10,
    help="average number of items per spot",
)
parser.add_argument(
    "--out-items",
    dest="out_items",
    type=argparse.FileType("w"),
    help="output items file name",
)


def main():
    args = parser.parse_args(sys.argv[1:])
    rng = faker.Faker()
    for s in range(args.nspots):
        args.out_spots.write(f"{json.dumps(gen_spot(rng, s))}\n")
    o = 10 ** math.ceil(math.log10(args.nspots))
    for s in range(args.nspots):
        for i in range(max(1, random.randint(args.nitems // 2, args.nitems * 2))):
            args.out_items.write(f"{json.dumps(gen_item(rng, s, s * o + i))}\n")


def gen_spot(rng: faker.Faker, spot: int) -> dict:
    lat, lng, city, country, state = rng.local_latlng()
    return {
        "id": spot,
        "name": rng.company(),
        "description": rng.sentence(nb_words=20),
        "raiting": rng.pyfloat(min_value=1, max_value=10),
        "country": country,
        "state": state,
        "city": city,
        "lat": lat,
        "lng": lng,
    }


def gen_item(rng: faker.Faker, spot: int, item: int) -> dict:
    return {
        "id": item,
        "spot": spot,
        "raiting": rng.pyfloat(min_value=1, max_value=10),
        "name": rng.sentence(nb_words=3),
        "description": rng.sentence(nb_words=20),
    }


if __name__ == "__main__":
    main()
