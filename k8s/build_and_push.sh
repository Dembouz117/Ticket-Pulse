docker build --platform linux/amd64 -t tigercxx/cs203-auth:1.0 ../auth
docker build --platform linux/amd64 -t tigercxx/cs203-ticketing:1.0 ../ticketing
docker build --platform linux/amd64 -t tigercxx/cs203-payment:1.0 ../payment
docker build --platform linux/amd64 -t tigercxx/cs203-queue:1.0 ../queue
docker build --platform linux/amd64 \
--build-arg NEXT_PUBLIC_STRIPE_PUBLISHABLE_KEY="pk_test_51Kia4HFtg42uLJvvAAj23WtQbTAzT8ZFK9z1uhIsgK848qUvMxCIerOmOC1ldbyRi19lzI7nq6uvGdxmu6Q7bJmV00sMf1m1da" \
--build-arg NEXT_PUBLIC_AUTH_API_URL="https://sg1.biddlr.com/auth" \
--build-arg NEXT_PUBLIC_TICKETING_API_URL="https://sg1.biddlr.com/ticketing" \
--build-arg NEXT_PUBLIC_PAYMENT_API_URL="https://sg1.biddlr.com/payment" \
--build-arg NEXT_PUBLIC_QUEUE_API_URL="https://sg1.biddlr.com/queue" \
--build-arg NEXT_PUBLIC_QUEUE_WEBSOCKET_URL="wss://sg1.biddlr.com/queue/queue" \
--build-arg NEXT_PUBLIC_FRONTEND_URL="https://sg1.biddlr.com" \
-t tigercxx/cs203-frontend:1.0 ../frontend


docker push tigercxx/cs203-auth:1.0
docker push tigercxx/cs203-ticketing:1.0
docker push tigercxx/cs203-payment:1.0
docker push tigercxx/cs203-queue:1.0
docker push tigercxx/cs203-frontend:1.0
