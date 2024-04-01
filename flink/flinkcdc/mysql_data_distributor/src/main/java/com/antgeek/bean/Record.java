package com.antgeek.bean;

import lombok.Data;
import lombok.ToString;

import java.io.Serializable;

@Data
@ToString
public class Record implements Serializable {
    private static final long serialVersionUID = 1L;
    private Long id;
    private String userId;
    private String record;
}
