import argparse
import io
import json
import math
import os
import random
import sys
import time

import faker
import jwt

parser = argparse.ArgumentParser(prog="datagen")

subparsers = parser.add_subparsers(dest="command")

payload_parser = subparsers.add_parser("payload")
payload_parser.add_argument(
    "-n",
    dest="nspots",
    type=int,
    default=1,
    help="number of spots",
)
payload_parser.add_argument(
    "-i",
    dest="nitems",
    type=int,
    default=10,
    help="average number of items per spot",
)
payload_parser.add_argument(
    "-t",
    dest="nthreads",
    type=int,
    default=2,
    help="number of threads",
)
outdir_arg = payload_parser.add_argument(
    "-o",
    dest="outdir",
    type=str,
    default=".",
    help="output directory path",
)

jwt_parser = subparsers.add_parser("jwt")
jwt_parser.add_argument(
    "-o",
    dest="out",
    type=argparse.FileType("w"),
    required=True,
    help="output file name",
)
jwt_parser.add_argument(
    "-k",
    dest="key",
    default=os.getenv("JWT_KEY"),
    type=str,
    required=True,
    help="signing key",
)


def main():
    args = parser.parse_args(sys.argv[1:])
    try:
        match args.command:
            case "payload":
                gen_payload(args)
            case "jwt":
                gen_jwt(args)
            case _:
                parser.print_help(sys.stderr)
    except argparse.ArgumentError as ex:
        sys.stderr.write(f"{ex}\n")


def gen_payload(args) -> None:
    rng = faker.Faker()
    if not os.path.exists(args.outdir):
        raise argparse.ArgumentError(
            outdir_arg,
            "output directory does not exist",
        )
    s = 0
    for t in range(1, args.nthreads + 1):
        with open(os.path.join(args.outdir, f"spots.{t}.jsonl"), "w") as fd:
            for _ in range(args.nspots//args.nthreads):
                s += 1
                fd.write(f"{json.dumps(gen_spot(rng, s))}\n")
    # for s in range(args.nspots):
    #     args.out_spots.write(f"{json.dumps(gen_spot(rng, s))}\n")
    # o = 10 ** math.ceil(math.log10(args.nspots))
    # for s in range(args.nspots):
    #     for i in range(max(1, random.randint(args.nitems // 2, args.nitems * 2))):
    #         args.out_items.write(f"{json.dumps(gen_item(rng, s, s * o + i))}\n")


def gen_jwt(args) -> None:
    token = jwt.encode(
        payload={
            "sub": "1",
            "name": "John Doe",
            "iat": int(time.time()),
            "roles": 7,
        },
        key=args.key,
        algorithm="HS256",
    )
    args.out.write(token)


def gen_spot(rng: faker.Faker, spot: int) -> dict:
    coorinates = rng.local_latlng()
    assert coorinates, "failed generate coorinates"
    lat, lng, city, country, state = coorinates
    return {
        "id": spot,
        "name": rng.company(),
        "description": rng.sentence(nb_words=20),
        "raiting": rng.pyfloat(min_value=1, max_value=10),
        "country": country,
        "state": state,
        "city": city,
        "lat": float(lat),
        "lng": float(lng),
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
