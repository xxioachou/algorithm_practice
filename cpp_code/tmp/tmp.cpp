#include <iostream>
#include <memory>
#include <stdexcept>
#include <algorithm>
#include <vector>

template <typename T>
class Vector {
public:
    Vector () : size_(0), capacity_(1), data_(std::make_unique<T[]>(1)) {}
    ~Vector() = default;

    Vector(int size, const T &obj) {
        if (size < 0) {
            throw std::invalid_argument("size cannot be negative.");
        }

        size_ = size;
        capacity_ = std::max(size, 1);
        data_ = std::make_unique<T[]>(capacity_);
        for (int i = 0; i < size; i ++) {
            data_[i] = obj;
        }
    }
    Vector(Vector &&other) noexcept : size_(other.size_), capacity_(other.capacity_), data_(std::move(other.data_)) {
        other.size_ = 0;
        other.capacity_ = 0;
    }   
    Vector& operator=(Vector &&other) noexcept {
        if (this != &other) {
            data_ = std::move(other.data_);
            size_ = other.size_;
            capacity_ = other.capacity_;
            other.size_ = 0;
            other.capacity_ = 0;
        }
    }

    /*
        在数组末尾插入一个数，如果 size >= capacity_ 则调用 reserve() 扩充数组
    */
    void push_back(const T &obj) {
        if (size_ >= capacity_) {
            reserve(capacity_ * 2);
        }
        
        data_[size_ ++] = obj;
    }

    /*
        扩充数组容量
    */
    void reserve(int new_capacity) {
        if (new_capacity <= 0) {
            throw std::invalid_argument("new capacity cannot be negative.");
        }
        if (new_capacity <= capacity_) {
            return;
        }
        std::cout << "reserve capacity_ " << capacity_ << " to new_capacity " << new_capacity << std::endl;

        auto new_data = std::make_unique<T[]>(new_capacity);
        for (int i = 0; i < size_; i ++) {
            new_data[i] = data_[i];
        }
        data_ = std::move(new_data);
        capacity_ = new_capacity;
    }

    /*
        把数组的 size 变为 new_size
            如果 size < new_size，填充默认构造的 T；
            如果 size >= new_size，截断数组；
    */
    void resize(int new_size) {
        if (new_size < 0) {
            throw std::invalid_argument("new_size cannot be negative.");
        }

        int new_capacity = capacity_;
        while (new_size > new_capacity) {
            new_capacity *= 2;
        }
        
        auto new_data = std::make_unique<T[]>(new_capacity);
        int cnt = std::min(size_, new_size);
        for (int i = 0; i < cnt; i ++) {
            new_data[i] = data_[i];
        }
        for (int i = cnt; i < new_size; i ++) {
            new_data[i] = T();
        }
        
        size_ = new_size;
        capacity_ = new_capacity;
        data_ = std::move(new_data);
    }

    int capacity() const { return capacity_; }

    int size() const { return size_; }

    void print() const {
        std::cout << "[";
        for (int i = 0; i < size_; i ++) {
            std::cout << data_[i] << ":"[i == size_ - 1];
        }
        std::cout << "]" << std::endl;
    }

private:
    std::unique_ptr<T[]> data_;
    int capacity_;
    int size_;
};

int main() {
    {
        Vector<int> v2(5, 666);
        v2.print();
        std::cout << "v2.size(): " << v2.size() << std::endl;
    }

    {
        Vector<int> v1;
        for (int i = 0; i < 20; i ++) {
            v1.push_back(i);
        }
        v1.print();
        std::cout << "v1.size(): " << v1.size() << std::endl;

        v1.resize(10);
        v1.print();
        std::cout << "v1.size(): " << v1.size() << std::endl;

        v1.resize(20);
        v1.print();
        std::cout << "v1.size(): " << v1.size() << std::endl;

        v1.resize(0);
        v1.print();
        std::cout << "v1.size(): " << v1.size() << std::endl;

        // v1.resize(-1);
        // 抛出异常
    }

    {
        Vector<std::string> v1;
        v1.push_back("Hello");
        v1.push_back(",");
        v1.push_back("CPP");
        v1.print();

        auto v2 = std::move(v1);
        v2.print();

        v2.push_back("!");
        v2.print();
    }

    // std::vector<int> v;
    return 0;
}