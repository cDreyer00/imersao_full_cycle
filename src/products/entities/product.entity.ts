import { Entity, PrimaryGeneratedColumn, Column } from "typeorm";

@Entity('products')
export class Product {
   @PrimaryGeneratedColumn('uuid')
   id:          string;

   @Column()
   name:        string;

   @Column({ type: 'text'})
   description: string;

   @Column()
   price:       number;

   @Column()
   image_url:   string;

}
