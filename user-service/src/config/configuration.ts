import ConfigInterface from 'src/config/config.interface';

export default (): ConfigInterface => ({
  env: process.env.NODE_ENV || 'development',
  port: parseInt(process.env.PORT || '3000', 10),
  mongo_db: process.env.MONGODB_URI || 'mongodb://localhost:27017/user-service',
  swagger: {
    username: process.env.SWAGGER_USERNAME || '',
    password: process.env.SWAGGER || ''
  },
  redis_db: process.env.REDIS_URI || 'redis://localhost:6379',
  nats_uri: process.env.NATS_URI || 'nats://localhost:4222',
  nats_user: process.env.NATS_USER || '',
  nats_pass: process.env.NATS_PASS || '',
  rabbitmq_url: process.env.RABBITMQ_URL || 'amqp://localhost:5672',
  bucket_name: process.env.AWS_BUCKET_NAME || '',
  cloudfront_public: process.env.AWS_CLOUDFRONT_PUBLIC_KEY || '',
  cloudfront_private: process.env.AWS_CLOUDFRONT_PRIVATE_KEY || '',
  cloudfront_url: process.env.AWS_CLOUDFRONT_URL || '',
  aws_key: process.env.AWS_ACCESS_KEY_ID || '',
  aws_secret: process.env.AWS_SECRET_ACCESS_KEY || '',
  from_address: process.env.AWS_SENDER_EMAIL || 'datnguyen03011985@gmail.com',
  google: {
    oauth_google_id: process.env.OAUTH_GOOGLE_ID || '',
    oauth_google_secret: process.env.OAUTH_GOOGLE_SECRET || '',
    oauth_google_callback: process.env.OAUTH_GOOGLE_REDIRECT_URL || ''
  },
  facebook: {
    oauth_facebook_id: process.env.OAUTH_FACEBOOK_ID || '',
    oauth_facebook_secret: process.env.OAUTH_FACEBOOK_SECRET || '',
    oauth_facebook_callback: process.env.OAUTH_FACEBOOK_REDIRECT_URL || ''
  }
});
