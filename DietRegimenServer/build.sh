RUN_NAME="DietRegimen"

mkdir -p output/bin output/conf
cp script/* output
chmod +x output/bootstrap.sh
chmod +x output/bootstrap_staging.sh
cp conf/* output/conf

find conf/ -type f ! -name *_local.* | xargs -I{} cp {} output/conf/
go build -o output/bin/$RUN_NAME
