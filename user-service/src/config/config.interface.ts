export default interface ConfigInterface {
  env: string;
  port: number;
  mongo_db: string;
  swagger: {
    username: string;
    password: string;
  };
  google: {
    oauth_google_id: string;
    oauth_google_secret: string;
    oauth_google_callback: string;
  };
  facebook: {
    oauth_facebook_id: string;
    oauth_facebook_secret: string;
    oauth_facebook_callback: string;
  };
  redis_db: string;
  nats_uri: string;
  nats_user: string;
  nats_pass: string;
  rabbitmq_url: string;
  aws_key: string;
  aws_secret: string;
  bucket_name: string;
  cloudfront_public: string;
  cloudfront_private: string;
  cloudfront_url: string;
  from_address: string;
}
